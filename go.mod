module github.com/meshplus/pier-client-ethereum

go 1.13

require (
	github.com/Rican7/retry v0.1.0
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/ethereum/go-ethereum v1.10.7
	github.com/fatih/color v1.9.0
	github.com/gobuffalo/packd v1.0.0
	github.com/gobuffalo/packr v1.30.1
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	github.com/hashicorp/go-plugin v1.3.0
	github.com/meshplus/bitxhub-model v1.2.1-0.20220803022708-9ab7a71abdbf
	github.com/meshplus/pier v1.24.1-0.20220803023357-8533944f0d08
	github.com/spf13/viper v1.7.1
	github.com/urfave/cli v1.22.1
)

replace github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.10.2

replace github.com/binance-chain/tss-lib => github.com/dawn-to-dusk/tss-lib v1.3.3-0.20220330081758-f404e10a1268
