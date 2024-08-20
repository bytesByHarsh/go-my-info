package router

import (
	"github.com/bytesByHarsh/go-my-info/handler"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	v1 := app.Group("/v1", logger.New())
	v1.Get("/", handler.Hello)

	// User
	user := v1.Group("/user")
	user.Post("/", handler.CreateUser)
}
