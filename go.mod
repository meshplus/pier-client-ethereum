module github.com/meshplus/pier-client-ethereum

go 1.13

require (
	github.com/Rican7/retry v0.1.0
	github.com/cloudflare/cfssl v1.4.1
	github.com/ethereum/go-ethereum v1.9.13
	github.com/meshplus/bitxhub-kit v1.0.1-0.20200727075316-ea098a3c3411
	github.com/meshplus/bitxhub-model v1.0.0-rc4.0.20200608065824-2fbc63639e92
	github.com/meshplus/pier v1.0.0-rc4
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.5.0
	github.com/spf13/viper v1.6.1
)

replace golang.org/x/sys => golang.org/x/sys v0.0.0-20200509044756-6aff5f38e54f

replace golang.org/x/text => golang.org/x/text v0.3.0

replace golang.org/x/net => golang.org/x/net v0.0.0-20200202094626-16171245cfb2

replace golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200311171314-f7b00557c8c4

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.2

replace github.com/sirupsen/logrus => github.com/sirupsen/logrus v1.5.0

replace github.com/meshplus/pier => ../pier

replace github.com/spf13/afero => github.com/spf13/afero v1.1.2

replace github.com/spf13/pflag => github.com/spf13/pflag v1.0.5

replace google.golang.org/grpc => google.golang.org/grpc v1.27.1

replace gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.7
