package utils

import (
	"github.com/google/wire"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type Utils struct {
	Logger Logger
	Tracer *sdktrace.TracerProvider
	Config *Config
}

func NewUtils(l Logger, t *sdktrace.TracerProvider, cfg *Config) (*Utils, error) {
	return &Utils{
		Logger: l,
		Tracer: t,
		Config: cfg,
	}, nil
}

var UtilsSet = wire.NewSet(
	NewLogger, NewTracer, NewConfig,
)
