package main

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	configName = "ethereum.toml"
)

type Config struct {
	Ether Ether `toml:"ether" json:"ether"`
}

type Ether struct {
	Addr            string `toml:"addr" json:"addr"`
	Name            string `toml:"name" json:"name"`
	ContractAddress string `mapstructure:"contract_address" json:"contract_address"`
	AbiPath         string `mapstructure:"abi_path" json:"abi_path"`
	KeyPath         string `mapstructure:"key_path" json:"key_path"`
}

func defaultConfig() *Config {
	return &Config{
		Ether: Ether{
			Addr:            "https://mainnet.infura.io",
			Name:            "Ethereum",
			ContractAddress: "0x1049a0bc31bb746b56b2c10d81644ab7579eb45b",
			AbiPath:         "broker.abi",
			KeyPath:         "~/.pier/ether/key.json",
		},
	}
}

func UnmarshalConfig(configRoot string) (*Config, error) {
	viper.SetConfigFile(filepath.Join(configRoot, configName))
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("ETHER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := defaultConfig()

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
