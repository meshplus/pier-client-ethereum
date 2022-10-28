SHELL := /bin/bash
CURRENT_PATH = $(shell pwd)
DISTRO = $(shell uname)
CURRENT_TAG =$(shell git describe --abbrev=0 --tags)

GO  = GO111MODULE=on go

ifeq (docker,$(firstword $(MAKECMDGOALS)))
  DOCKER_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(DOCKER_ARGS):;@:)
endif

help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## make eth: build ethereum client plugin
eth:
	@packr2
	mkdir -p build
	$(GO) build -o build/eth-client ./*.go

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

