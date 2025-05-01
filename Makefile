# ====================================================================================
# Project Specific Globals
# ------------------------------------------------------------------------------------
#
# - It's assumed the $(name) is the same literal as the compiled binary.
# - Override the defaults if not available in a pipeline's environment variables.
#
# - Default GitHub environment variables: https://docs.github.com/en/actions/writing-workflows/choosing-what-your-workflow-does/store-information-in-variables#default-environment-variables

name := example
ifdef CI_PROJECT_NAME
    override name = $(CI_PROJECT_NAME)
endif

homebrew-tap := example-organization/example-cli
ifdef HOMEBREW_TAP
    override homebrew-tap = $(HOMEBREW_TAP)
endif

homebrew-tap-repository := gitlab.com:example-organization/group-1/group-2/homebrew-taps.git
ifdef HOMEBREW_TAP_REPOSITORY
    override homebrew-tap-repository = $(HOMEBREW_TAP_REPOSITORY)
endif

type := patch
ifdef RELEASE_TYPE
    override type = RELEASE_TYPE
endif

ifeq (,$(shell go env GOBIN))
	GOBIN=$(shell go env GOPATH)/bin
else
	GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
SHELL = /usr/bin/env bash -o pipefail

.SHELLFLAGS = -ec

# ====================================================================================
# Colors
# ------------------------------------------------------------------------------------

black        := $(shell printf "\033[30m")
black-bold   := $(shell printf "\033[30;1m")
red          := $(shell printf "\033[31m")
red-bold     := $(shell printf "\033[31;1m")
green        := $(shell printf "\033[32m")
green-bold   := $(shell printf "\033[32;1m")
yellow       := $(shell printf "\033[33m")
yellow-bold  := $(shell printf "\033[33;1m")
blue         := $(shell printf "\033[34m")
blue-bold    := $(shell printf "\033[34;1m")
magenta      := $(shell printf "\033[35m")
magenta-bold := $(shell printf "\033[35;1m")
cyan         := $(shell printf "\033[36m")
cyan-bold    := $(shell printf "\033[36;1m")
white        := $(shell printf "\033[37m")
white-bold   := $(shell printf "\033[37;1m")
reset        := $(shell printf "\033[0m")

# ====================================================================================
# Logger
# ------------------------------------------------------------------------------------

time-long	= $(date +%Y-%m-%d' '%H:%M:%S)
time-short	= $(date +%H:%M:%S)
time		= $(time-short)

information	= echo $(time) $(green)[ INFO ]$(reset)
debug	= echo $(time) $(blue)[ DEBUG ]$(reset)
warning	= echo $(time) $(yellow)[ WARNING ]$(reset)
exception		= echo $(time) $(red)[ ERROR ]$(reset)
complete		= echo $(time) $(white)[ COMPLETE ]$(reset)
fail	= (echo $(time) $(red)[ FAILURE ]$(reset) && false)

# ====================================================================================
# Utility Command(s)
# ------------------------------------------------------------------------------------

url = $(shell git config --get remote.origin.url | sed -r 's/.*(\@|\/\/)(.*)(\:|\/)([^:\/]*)\/([^\/\.]*)\.git/https:\/\/\2\/\4\/\5/')

repository = $(shell basename -s .git $(shell git config --get remote.origin.url))
organization = $(shell git remote -v | grep "(fetch)" | sed 's/.*\/\([^ ]*\)\/.*/\1/')
package = $(shell git remote -v | grep "(fetch)" | sed 's/^origin[[:space:]]*//; s/[[:space:]]*(fetch)$$//' | sed 's/https:\/\///; s/git@//; s/\.git$$//; s/:/\//')

version = $(shell [ -f VERSION ] && head VERSION || echo "0.0.0")

major      		= $(shell echo $(version) | sed "s/^\([0-9]*\).*/\1/")
minor      		= $(shell echo $(version) | sed "s/[0-9]*\.\([0-9]*\).*/\1/")
patch      		= $(shell echo $(version) | sed "s/[0-9]*\.[0-9]*\.\([0-9]*\).*/\1/")

zero = $(shell printf "%s" "0")

major-upgrade 	= $(shell expr $(major) + 1).$(zero).$(zero)
minor-upgrade 	= $(major).$(shell expr $(minor) + 1).$(zero)
patch-upgrade 	= $(major).$(minor).$(shell expr $(patch) + 1)

dirty = $(shell git diff --quiet)
dirty-contents 			= $(shell git diff --shortstat 2>/dev/null 2>/dev/null | tail -n1)

# ====================================================================================
# Build Command(s)
# ------------------------------------------------------------------------------------

compile = go build --mod "vendor" --ldflags "-s -w -X=main.version=$(tag) -X=main.date=$(shell date +%Y-%m-%d:%H-%M-%S) -X=main.source=false" -o "./build/$(name)-$(GOOS)-$(GOARCH)/$(name)"
compile-windows = go build --mod "vendor" --ldflags "-s -w -X=main.version=$(tag) -X=main.date=$(shell date +%Y-%m-%d:%H-%M-%S) -X=main.source=false" -o "./build/$(name)-$(GOOS)-$(GOARCH)/$(name).exe"

archive = tar -czvf "$(name)-$(GOOS)-$(GOARCH).tar.gz" -C "./build/$(name)-$(GOOS)-$(GOARCH)" .
archive-windows = cd "./build/$(name)-$(GOOS)-$(GOARCH)" && zip -r "../../$(name)-$(GOOS)-$(GOARCH).zip" "." && cd -

distribute = mkdir -p distribution && mv *.tar.gz distribution
distribute-windows = mkdir -p distribution && mv *.zip distribution

# ====================================================================================
# Default
# ------------------------------------------------------------------------------------

all :: prepare

# ====================================================================================
# Package-Specific Target(s)
# ------------------------------------------------------------------------------------

.PHONY: prepare
prepare:
	@$(information) Tidying and Reformatting Package
	@go mod tidy && go mod vendor
	@goimports -w .
	@go fmt ./...
	@$(complete)


.PHONY: test
test: prepare
	@$(information) Running Test Suite
	@go test -v --fullpath --cover --tags local ./...
	@$(complete)