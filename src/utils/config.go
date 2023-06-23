package utils

import (
	"github.com/caarlos0/env/v8"
)

type Config struct {
	Dev         bool   `env:"DEV" envDefault:"true"`
	Debug       bool   `env:"DEBUG" envDefault:"true"`
	LogLevel    string `env:"LOG_LEVEL" envDefault:"info"`
	TraceEnable bool   `env:"TRACE_ENABLE" envDefault:"false"`
	TraceUrl    string `env:"TRACE_URL"`
}

// type EnvConfig struct{}

func NewConfig(filePath *string) (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
