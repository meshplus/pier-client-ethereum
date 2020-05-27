SHELL := /bin/bash
CURRENT_PATH = $(shell pwd)

GO  = GO111MODULE=on go

help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## make eth: build ethereum client plugin
eth:
	mkdir -p build
	$(GO) build --buildmode=plugin -o build/eth-client.so ./*.go

## make linter: Run golanci-lint
linter:
	golangci-lint run -E goimports --skip-dirs-use-default -D staticcheck

