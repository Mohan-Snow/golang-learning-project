package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port             int    `env:"PORT" envDefault:"8080"`
	ExternalApiToken string `env:"EXTERNAL_API_TOKEN,notEmpty,unset,file"`
	Username         string
	Password         string
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
