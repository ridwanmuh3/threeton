package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"threeton-starter/internal/delivery/http/handler"
	"threeton-starter/internal/delivery/http/route"
	"threeton-starter/internal/repository"
	"threeton-starter/internal/service"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *zap.SugaredLogger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// repositories
	testRepository := repository.NewTestRepository()

	// services
	testService := service.NewTestService(config.DB, config.Validate, config.Log, testRepository)

	// handler
	testHandler := handler.NewTestHandler(config.Log, testService)

	routeConfig := &route.RouteConfig{
		App:         config.App,
		TestHandler: testHandler,
		Log:         config.Log,
	}

	routeConfig.Setup()
}
