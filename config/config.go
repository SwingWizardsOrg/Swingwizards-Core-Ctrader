package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Trader struct {
		ClientId     string `yaml:"client_id"`
		AccountId    string `yaml:"account_id"`
		ClientSecret string `yaml:"client_secret"`
		AccessToken  string `yaml:"access_token"`
	} `yaml:"account"`

	Server struct {
		Endpoint string `yaml:"endpoint"`
		Port     int32  `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig() (*Config, error) {
	f, err := os.Open("./config.yml")
	if err != nil {
		return nil, err
	}
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
