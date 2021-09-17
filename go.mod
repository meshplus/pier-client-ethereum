module github.com/meshplus/pier-client-ethereum

go 1.13

require (
	github.com/Rican7/retry v0.1.0
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/ethereum/go-ethereum v1.10.4
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	github.com/hashicorp/go-plugin v1.3.0
	github.com/meshplus/bitxhub-model v1.2.1-0.20210918014850-45a1a094b597
	github.com/meshplus/pier v1.12.1-0.20210917172218-59d1246c8859
	github.com/spf13/viper v1.7.1
)

replace github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.10.2
