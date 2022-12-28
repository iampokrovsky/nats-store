package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		NATS `yaml:"nats"`
		HTTP `yaml:"http"`
	}

	NATS struct {
		ClusterID string `yaml:"cluster_id"`
		ClientID  string `yaml:"client_id"`
		Subject   string `yaml:"subject"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
