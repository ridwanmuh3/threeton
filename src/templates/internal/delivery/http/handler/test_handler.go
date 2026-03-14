package handler

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"threeton-starter/internal/exception"
	"threeton-starter/internal/model"
	"threeton-starter/internal/service"
)

type TestHandler struct {
	log         *zap.SugaredLogger
	testService *service.TestService
}

func NewTestHandler(log *zap.SugaredLogger, testService *service.TestService) *TestHandler {
	return &TestHandler{
		log:         log,
		testService: testService,
	}
}

func (h *TestHandler) SayHello(c fiber.Ctx) error {
	name := c.Params("name", "John Doe")
	msg, err := h.testService.SayHello(c.Context(), &model.SayHelloRequest{Name: name})
	if err != nil {
		h.log.Warnf("failed to say hello: %v", err)
		return exception.InternalServerError
	}

	return c.JSON(model.Response[string]{
		Status:  fiber.StatusOK,
		Message: msg,
	})
}
