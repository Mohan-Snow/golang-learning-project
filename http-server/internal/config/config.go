package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port                 string `env:"PORT" envDefault:"8080"`
	GenerationQueryParam string `env:"GENERATION_QUERY_PARAM,notEmpty,unset,file"`
	DbConnection         string `env:"DB_CONNECTION,notEmpty,unset,file"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
