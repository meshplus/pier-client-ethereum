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

# build with version infos
VERSION_DIR = github.com/meshplus/${APP_NAME}/main
GOLDFLAGS += -X "main.BuildDate=${BUILD_DATE}"
GOLDFLAGS += -X "main.CurrentCommit=${GIT_COMMIT}"
GOLDFLAGS += -X "main.CurrentBranch=${GIT_BRANCH}"
GOLDFLAGS += -X "main.CurrentVersion=${APP_VERSION}"

ifndef (${TAG})
  TAG = latest
endif

GREEN=\033[0;32m
NC=\033[0m

help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

prepare:
	cd scripts && bash prepare.sh

## make test-coverage: Test project with cover
test-coverage: prepare
	@go test -short -coverprofile cover.out -covermode=atomic ${TEST_PKGS}
	@cat cover.out >> coverage.txt

## make eth: build ethereum client plugin
eth:
	@packr2
	mkdir -p build
	$(GO) build -o eth-client -ldflags '${GOLDFLAGS}' *.go
	@mv eth-client build/eth-client
	@printf "${GREEN}Build eth-client successfully!${NC}\n"

## make build-docker: docker build the project
build-docker:
	docker build -t meshplus/pier-ethereum:${TAG} .
	@printf "${GREEN}Build images meshplus/pier-ethereum:${TAG} successfully!${NC}\n".

release-binary:
	mkdir -p build
	$(GO) build -o build/eth-client-${CURRENT_TAG}-${DISTRO} ./*.go

## make linter: Run golanci-lint
linter:
	golangci-lint run -E goimports --skip-dirs-use-default -D staticcheck

