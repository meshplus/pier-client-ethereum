module github.com/meshplus/pier-client-ethereum

go 1.13

require (
	github.com/Rican7/retry v0.1.0
	github.com/cloudflare/cfssl v1.4.1
	github.com/ethereum/go-ethereum v1.10.4
	github.com/hashicorp/go-hclog v0.0.0-20180709165350-ff2cf002a8dd
	github.com/hashicorp/go-plugin v1.3.0
	github.com/meshplus/bitxhub-core v1.3.1-0.20210708054838-00de516e0ffd
	github.com/meshplus/bitxhub-model v1.2.1-0.20210701090843-8709b8dc88a6
	github.com/meshplus/bitxid v0.0.0-20210412025850-e0eaf0f9063a
	github.com/meshplus/pier v1.7.1-0.20210701092509-5ee894fbed5b
	github.com/pkg/errors v0.9.1
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.7.0
)

replace github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.10.2
