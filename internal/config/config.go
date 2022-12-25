package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Server struct {
		Host string `envconfig:"SERVER_HOST" default:"127.0.0.1:9000"`
	}
	DataBase struct {
		Host     string `envconfig:"DB_HOST" default:"127.0.0.1"`
		Database string `envconfig:"DB_DATABASE" default:"postgres"`
		User     string `envconfig:"DB_USER" default:"maui"`
		Password string `envconfig:"DB_PASSWORD" default:"maui"`
	}
}

func Parse() (*Config, error) {
	var cfg = &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
