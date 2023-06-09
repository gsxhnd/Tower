package utils

import "github.com/caarlos0/env/v8"

type Config struct {
	Dev      bool   `env:"DEV"`
	Debug    bool   `env:"DEBUG"`
	LogLevel string `env:"LOG_LEVEL"`
}

// type EnvConfig struct{}

func NewConfig(filePath *string) (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
