module github.com/meshplus/pier-client-ethereum

go 1.13

require (
	github.com/Rican7/retry v0.1.0
	github.com/cloudflare/cfssl v1.4.1
	github.com/ethereum/go-ethereum v1.10.2
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	github.com/hashicorp/go-plugin v1.3.0
	github.com/meshplus/bitxhub-core v1.3.1-0.20210611011450-ca11d623d8cc
	github.com/meshplus/bitxhub-model v1.2.1-0.20210616124147-a2470bd3d55d
	github.com/meshplus/bitxid v0.0.0-20210412025850-e0eaf0f9063a
	github.com/meshplus/pier v1.7.1-0.20210524093640-1337e0a53318
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.0
)

replace github.com/meshplus/pier => ../pier
