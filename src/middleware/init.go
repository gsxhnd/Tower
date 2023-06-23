package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gsxhnd/go-api-template/src/utils"
)

type Middlewarer interface {
	RequestLog() gin.HandlerFunc
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
