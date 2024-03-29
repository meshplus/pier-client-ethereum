package main

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	configName = "ethereum.toml"
	directMode = "direct"
	relayMode  = "relay"
)

type Config struct {
	Ether Ether `toml:"ether" json:"ether"`
}

type Ether struct {
	Addr            string `toml:"addr" json:"addr"`
	Name            string `toml:"name" json:"name"`
	ContractAddress string `mapstructure:"contract_address" json:"contract_address"`
	KeyPath         string `mapstructure:"key_path" json:"key_path"`
	Password        string `toml:"password" json:"password"`
	MinConfirm      uint64 `mapstructure:"min_confirm" json:"min_confirm"`
	TimeoutHeight   uint64 `mapstructure:"timeout_height" json:"timeout_height"`
	TimeoutPeriod   uint64 `mapstructure:"timeout_period" json:"timeout_period"`
	ChainID         string `mapstructure:"chain_id" json:"chain_id"`
	OffChainAddr    string `mapstructure:"offchain_addr" json:"offchain_addr"`
	OffChainPath    string `mapstructure:"offchain_path" json:"offchain_path"`
}

func defaultConfig() *Config {
	return &Config{
		Ether: Ether{
			Addr:            "https://mainnet.infura.io",
			Name:            "Ethereum",
			ContractAddress: "0xD3880ea40670eD51C3e3C0ea089fDbDc9e3FBBb4",
			KeyPath:         "account.key",
			Password:        "",
			MinConfirm:      15,
			TimeoutHeight:   100,
			TimeoutPeriod:   60,
			OffChainAddr:    "",
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
