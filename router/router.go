package router

import (
	"trainee3/handlers"

	"trainee3/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func Set(handlers *handlers.Handlers) {
	app := fiber.New(fiber.Config{
		ErrorHandler: nil,
	})

	app.Use(logger.New())
	app.Post("/auth/signIn", handlers.Auth.SignIn)
	app.Get("/cache/refresh", handlers.Cache.Refresh)

	app.Listen(":3010")

}
