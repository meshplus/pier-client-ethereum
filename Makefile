SHELL := /bin/bash
CURRENT_PATH = $(shell pwd)
DISTRO = $(shell uname)
CURRENT_TAG =$(shell git describe --abbrev=0 --tags)

GO  = GO111MODULE=on go
APP_NAME = pier-client-ethereum
BUILD_DATE = $(shell date +%FT%T)
GIT_COMMIT = $(shell git log --pretty=format:'%h' -n 1)
GIT_BRANCH = $(shell git rev-parse --abbrev-ref HEAD)
ifeq (${GIT_BRANCH},HEAD)
  APP_VERSION = $(shell git describe --tags HEAD)
else
  APP_VERSION = dev
endif

# build with verison infos
VERSION_DIR = github.com/meshplus/${APP_NAME}/main
GOLDFLAGS += -X "main.BuildDate=${BUILD_DATE}"
GOLDFLAGS += -X "main.CurrentCommit=${GIT_COMMIT}"
GOLDFLAGS += -X "main.CurrentBranch=${GIT_BRANCH}"
GOLDFLAGS += -X "main.CurrentVersion=${APP_VERSION}"

ifeq (docker,$(firstword $(MAKECMDGOALS)))
  DOCKER_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(DOCKER_ARGS):;@:)
endif

GREEN=\033[0;32m
NC=\033[0m

help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## make eth: build ethereum client plugin
eth:
	mkdir -p build
	$(GO) build -ldflags '${GOLDFLAGS}' *.go
	@mv broker build/eth-client
	@printf "${GREEN}Build eth-client successfully!${NC}\n"

docker:
	mkdir -p build
	cd build && rm -rf pier && git clone https://github.com/meshplus/pier.git && cd pier && git checkout $(DOCKER_ARGS)
	cd ${CURRENT_PATH}
	docker build -t meshplus/pier-ethereum .

release-binary:
	mkdir -p build
	$(GO) build -o build/eth-client-${CURRENT_TAG}-${DISTRO} ./*.go

## make linter: Run golanci-lint
linter:
	golangci-lint run -E goimports --skip-dirs-use-default -D staticcheck

