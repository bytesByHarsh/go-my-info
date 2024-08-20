package main

import (
	"log"

	"github.com/bytesByHarsh/go-my-info/router"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: false,
		StrictRouting: false,
		ServerHeader:  "Fiber",
		AppName:       "My Info",
	})

	router.SetupRoutes(app)
	log.Fatalf("Closing Server: %v", app.Listen(":3000", fiber.ListenConfig{EnablePrefork: true}))
}
