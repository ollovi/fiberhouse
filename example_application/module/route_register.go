package module

import (
	"github.com/gofiber/fiber/v2"
	exampleApi "github.com/lamxy/fiberhouse/example_application/module/example-module/api"
	"github.com/lamxy/fiberhouse/frame"
)

// RegisterRouteHandlers 注册各业务模块的路由处理器
func RegisterRouteHandlers(ctx frame.ContextFramer, app fiber.Router) fiber.Router {
	// 注册example模块的路由处理器
	exampleApi.RegisterRouteHandlers(ctx, app)

	// TODO 注册更多业务模块路由处理器 ...

	return app
}
