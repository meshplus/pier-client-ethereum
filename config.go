package main

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const configName = "mock.toml"

type Config struct {
	Mock Mock `toml:"mock" json:"mock"`
}

type Mock struct {
	Name          string   `toml:"name" json:"name"`
	TimeoutHeight uint64   `mapstructure:"timeout_height" json:"timeout_height"`
	Port          string   `toml:"port" json:"port"`
	BxhId         string   `toml:"bxhId" json:"bxhId"`
	ChainId       string   `toml:"chainId" json:"chainId"`
	ServiceList   []string `toml:"serviceList" json:"serviceList"`
}

func defaultConfig() *Config {
	return &Config{
		Mock: Mock{
			Name:          "Ethereum",
			Port:          "8081",
			BxhId:         "1356",
			ChainId:       "testChain",
			TimeoutHeight: 100,
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
