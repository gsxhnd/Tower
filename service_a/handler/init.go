package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	RootHandler      RootHandler
	WebsocketHandler WebsocketHandler
}

var HandlerSet = wire.NewSet(
	NewRootHandler,
	NewWebsocketHandler,
	wire.Struct(new(Handler), "*"),
)
