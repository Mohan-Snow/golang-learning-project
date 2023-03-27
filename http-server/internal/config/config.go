package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port                 string `env:"PORT" envDefault:"8080"`
	GenerationQueryParam string `env:"GENERATION_QUERY_PARAM,notEmpty,unset,file"`
	DbPort               string `env:"DB_PORT" envDefault:"5432"`
	DbHost               string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DbUser               string `env:"DB_USERNAME" envDefault:"postgres"`
	DbPassword           string `env:"DB_PASSWORD" envDefault:"postgres"`
	DbName               string `env:"DB_NAME" envDefault:"postgres"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
