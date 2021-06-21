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
	Ether       Ether             `toml:"ether" json:"ether"`
	ContractABI map[string]string `mapstructure:"contract_abi" json:"contract_abi"`
}

type Ether struct {
	Addr            string `toml:"addr" json:"addr"`
	Name            string `toml:"name" json:"name"`
	ContractAddress string `mapstructure:"contract_address" json:"contract_address"`
	EscrowsAddress  string `mapstructure:"escrows_address" json:"escrows_address"`
	KeyPath         string `mapstructure:"key_path" json:"key_path"`
	Password        string `toml:"password" json:"password"`
	MinConfirm      uint64 `mapstructure:"min_confirm" json:"min_confirm"`
}

func defaultConfig() *Config {
	return &Config{
		Ether: Ether{
			Addr:            "https://mainnet.infura.io",
			Name:            "Ethereum",
			ContractAddress: "0xD3880ea40670eD51C3e3C0ea089fDbDc9e3FBBb4",
			EscrowsAddress:  "0x956Be099e5Add3d95aaB9D1a7Da5a40eB9d02528",
			KeyPath:         "account.key",
			Password:        "",
			MinConfirm:      5,
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
