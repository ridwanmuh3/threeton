package route

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"threeton-starter/internal/delivery/http/handler"
)

type RouteConfig struct {
	App         *fiber.App
	Log         *zap.SugaredLogger
	TestHandler *handler.TestHandler
}

func (c *RouteConfig) Setup() {
	c.App.Get("/:name", c.TestHandler.SayHello)
}
