package middleware

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/semconv/v1.13.0/httpconv"
	"go.opentelemetry.io/otel/trace"
)

func (m *middleware) RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// savedCtx := c.Request.Context()
		tc := propagation.TraceContext{}
		ctx := tc.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
		opts := []trace.SpanStartOption{
			trace.WithNewRoot(),
			trace.WithAttributes(httpconv.ServerRequest("", c.Request)...),
			trace.WithSpanKind(trace.SpanKindServer),
		}
		ctx, span := m.tracer.Tracer("go_service_trace").Start(ctx, c.FullPath(), opts...)
		defer func() {
			// c.Request = c.Request.WithContext(savedCtx)
			span.End()
		}()

		span.AddEvent("middleware req log")
		span.SetAttributes(attribute.String("client_ip", c.ClientIP()), attribute.String("path", c.Request.RequestURI))
		c.Request = c.Request.WithContext(ctx)

		var key = []interface{}{
			"http_method", c.Request.Method,
			"path", c.Request.RequestURI,
			"client_ip", c.ClientIP(),
			"status", c.Writer.Status(),
		}
		if m.tracer.Enable() {
			key = append(key,
				"trace_id", span.SpanContext().TraceID(),
				"span_id", span.SpanContext().SpanID(),
			)
		}

		c.Next()
		m.logger.Infow("", key...)
	}

}
