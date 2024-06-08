package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/go-api-template/utils"
)

type Middlewarer interface {
	RequestLog(ctxc *fiber.Ctx) error
	// Websocket(ctxc *fiber.Ctx) error
}
type middleware struct {
	logger utils.Logger
	tracer utils.Tracer
}

func NewMiddleware(l utils.Logger, t utils.Tracer) Middlewarer {
	return &middleware{
		logger: l,
		tracer: t,
	}
}
