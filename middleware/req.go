package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func (m *middleware) RequestLog(ctx *fiber.Ctx) error {
	// tc := propagation.TraceContext{}
	// traceCtx := tc.Extract(c.Request.Context(), propagation.HeaderCarrier(c.Request.Header))
	// opts := []trace.SpanStartOption{
	// 	trace.WithNewRoot(),
	// 	trace.WithAttributes(httpconv.ServerRequest("", ctx.Request...),
	// 	trace.WithSpanKind(trace.SpanKindServer),
	// }
	// ctx, span := m.tracer.Tracer("go_service_trace").Start(ctx, c.FullPath(), opts...)
	// defer func() {
	// 	// c.Request = c.Request.WithContext(savedCtx)
	// 	span.End()
	// }()

	// span.AddEvent("middleware req log")
	// span.SetAttributes(attribute.String("client_ip", c.ClientIP()), attribute.String("path", c.Request.RequestURI))
	// ctx.Request = ctx.Request.WithContext(ctx)

	// if m.tracer.Enable() {
	// 	key = append(key,
	// 		"trace_id", span.SpanContext().TraceID(),
	// 		"span_id", span.SpanContext().SpanID(),
	// 	)
	// }

	ctx.Next()
	var key = []interface{}{
		"method", ctx.Method(),
		"path", ctx.Path(),
		"client_id", ctx.IP(),
		"status", ctx.Response().StatusCode(),
	}
	m.logger.Infow("", key...)
	return nil

}
