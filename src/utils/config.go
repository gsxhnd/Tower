package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dev      bool
	Debug    bool
	LogLevel string
}

// type EnvConfig struct{}

func NewConfig(filePath *string) (*Config, error) {
	var c Config
	data, err := ioutil.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &c); err != nil {
		return &c, err
	} else {
		return &c, err
	}
}
