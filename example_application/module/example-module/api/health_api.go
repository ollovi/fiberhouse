package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lamxy/fiberhouse/example_application/module/constant"
	"github.com/lamxy/fiberhouse/example_application/module/example-module/service"
	"github.com/lamxy/fiberhouse/frame"
	"github.com/lamxy/fiberhouse/frame/response"
)

type HealthHandler struct {
	frame.ApiLocator
	Service *service.HealthService
}

func NewHealthHandler(ctx frame.ContextFramer, serv *service.HealthService) *HealthHandler {
	name := GetKeyHealthHandler()
	return &HealthHandler{
		ApiLocator: frame.NewApi(ctx).SetName(name),
		Service:    serv,
	}
}

func GetKeyHealthHandler(ns ...string) string {
	return frame.RegisterKeyName("HealthHandler", frame.GetNamespace([]string{constant.NameModuleExample}, ns...)...)
}

func (ha *HealthHandler) Liveness(c *fiber.Ctx) error {
	result := ha.Service.GetHealth()
	return c.Status(fiber.StatusOK).JSON(response.RespSuccess(result))
}
