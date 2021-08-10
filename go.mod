module github.com/meshplus/pier-client-ethereum

go 1.13

require (
	github.com/Rican7/retry v0.1.0
	github.com/cloudflare/cfssl v1.4.1
	github.com/ethereum/go-ethereum v1.10.4
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	github.com/hashicorp/go-plugin v1.3.0
	github.com/meshplus/bitxhub-model v1.2.1-0.20210809062857-4adfa90d51e6
	github.com/meshplus/pier v1.11.1-0.20210810164927-0fcb00af02c5
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.0
)

replace github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.10.2
