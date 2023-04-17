package utils

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	Dev   bool
	Debug bool
}

type EnvConfig struct{}

func NewConfig(filePath *string) (*Config, error) {
	var c Config
	err := yaml.Unmarshal([]byte{}, &c)
	return &c, err
}
