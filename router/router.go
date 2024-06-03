package router

import (
	"trainee3/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/tom-riddle-sr/logger/middleware/logger"
)

func Set(handlers *handlers.Handlers) {
	app := fiber.New(fiber.Config{
		ErrorHandler: nil,
	})

	app.Use(logger.GetLogger())
	app.Post("/auth/signIn", handlers.Auth.SignIn)
	app.Get("/cache/refresh", handlers.Cache.Refresh)

	app.Listen(":3010")

}
