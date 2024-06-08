package utils

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.18.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type trace struct {
	provider oteltrace.TracerProvider
	enable   bool
}
type Tracer interface {
	Tracer(name string, options ...oteltrace.TracerOption) oteltrace.Tracer
	SpanFromContext(ctx context.Context) oteltrace.Span
	Enable() bool
}

func NewTracer(cfg *Config) (Tracer, error) {
	var t = &trace{}
	if cfg.TraceEnable {
		t.enable = true
		ctx := context.Background()
		res, err := resource.New(ctx,
			resource.WithAttributes(
				// semconv.SchemaURL,
				semconv.ServiceName("api-service"),
				semconv.ServiceVersion(version),
				attribute.String("log_level", cfg.LogLevel),
				// attribute.String("hostname", cfg.Hostname),
			),
		)
		if err != nil {
			return nil, err
		}
		ctx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()
		conn, err := grpc.DialContext(ctx, cfg.TraceUrl,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
		)
		if err != nil {
			return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
		}

		traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, fmt.Errorf("failed to create trace exporter: %w", err)
		}
		bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
		t.provider = sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithSpanProcessor(bsp),
			sdktrace.WithResource(res),
		)
	} else {
		t.enable = false
		t.provider = oteltrace.NewNoopTracerProvider()
	}

	return t, nil
}

func (t *trace) Tracer(name string, options ...oteltrace.TracerOption) oteltrace.Tracer {
	return t.provider.Tracer(name, options...)
}

func (t *trace) SpanFromContext(ctx context.Context) oteltrace.Span {
	return oteltrace.SpanFromContext(ctx)
}

func (t *trace) Enable() bool {
	return t.enable
}
