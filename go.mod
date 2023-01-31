module github.com/meshplus/pier-client-ethereum

go 1.13

require (
	github.com/Rican7/retry v0.1.0
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/ethereum/go-ethereum v1.10.22
	github.com/fatih/color v1.9.0
	github.com/gobuffalo/packd v1.0.1
	github.com/gobuffalo/packr v1.30.1
	github.com/gobuffalo/packr/v2 v2.8.3
	github.com/hashicorp/go-hclog v0.14.1
	github.com/hashicorp/go-plugin v1.3.0
	github.com/meshplus/bitxhub-core v1.28.0
	github.com/meshplus/bitxhub-model v1.28.0
	github.com/meshplus/pier v1.24.1-0.20230119083935-a568b0398d3c
	github.com/spf13/viper v1.8.1
	github.com/urfave/cli v1.22.1
)

// replace github.com/ethereum/go-ethereum => github.com/ethereum/go-ethereum v1.10.2

replace github.com/binance-chain/tss-lib => github.com/dawn-to-dusk/tss-lib v1.3.3-0.20220330081758-f404e10a1268

replace github.com/libp2p/go-libp2p-core => github.com/libp2p/go-libp2p-core v0.5.6

replace google.golang.org/grpc => google.golang.org/grpc v1.33.0
