package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/go-api-template/service_a/service"
	"github.com/gsxhnd/go-api-template/utils"
)

type RootHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type rootHandle struct {
	logger    utils.Logger
	validator *validator.Validate
	svc       service.PingService
}

func NewRootHandler(l utils.Logger, svc service.PingService) RootHandler {
	return &rootHandle{
		logger: l,
	}
}

func (h *rootHandle) Ping(ctx *fiber.Ctx) error {
	return ctx.Status(200).SendString("pong")
}
