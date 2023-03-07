package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port                 string `env:"PORT" envDefault:"8080"`
	GenerationQueryParam string `env:"GENERATION_QUERY_PARAM,notEmpty,unset,file"`
	Username             string `env:"DB_USERNAME" envDefault:"test_username"`
	Password             string `env:"DB_PASSWORD" envDefault:"test_password"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
