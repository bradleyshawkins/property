package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	ConnectionString string `envconfig:"DATABASE_URL" default:"postgresql://postgres:password@localhost:5432/rent-user?sslmode=disable"`
}

func Load() (*Config, error) {

	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
