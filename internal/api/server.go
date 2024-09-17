package api

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/api/handlers"
	"go-ecommerce-app/internal/api/rest"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	db, err := gorm.

	rh := &rest.RestHandler{
		App: app,
	}

	setupRoutes(rh)

	app.Listen(config.ServerPort)
}

func setupRoutes(rh *rest.RestHandler) {
	handlers.SetupUserRoute(rh)
}
