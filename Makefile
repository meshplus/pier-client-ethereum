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

docker:
	mkdir -p build
	cd build && rm -rf pier && git clone https://github.com/meshplus/pier.git && cd pier && git checkout v1.0.0-rc4
	cd ${CURRENT_PATH}
	docker build -t pier-ethereum .

## make linter: Run golanci-lint
linter:
	golangci-lint run -E goimports --skip-dirs-use-default -D staticcheck

