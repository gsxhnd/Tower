package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Dev         bool   `yaml:"dev"`
	Port        string `yaml:"port"`
	WebEnable   bool   `yaml:"web_enable"`
	TraceEnable bool   `yaml:"trace_enable"`
	TraceUrl    string
	LogConfig   LogConfig `yaml:"log,omitempty"`
	LogLevel    string
}

type LogConfig struct {
	Level       string
	Filename    string `yaml:"filename"`
	MaxSize     int
	MaxAge      int
	MaxBackups  int
	TraceEnable bool
	TraceUrl    string
}

// TODO: need config path
func NewConfig() (*Config, error) {
	var cfg = Config{}

	file, err := os.ReadFile("")
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
