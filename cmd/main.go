package main

import (
	"fmt"
	"log"

	"github.com/bytesByHarsh/go-my-info/config"
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

	config.ReadEnvFile(".env")

	// log.Printf("Starting Server at: %v:%v", config.Cfg.SERVER_LINK, config.Cfg.SERVER_PORT)

	serverAddr := fmt.Sprintf("%v:%v", config.Cfg.SERVER_LINK, config.Cfg.SERVER_PORT)

	router.SetupRoutes(app)
	log.Fatalf("Closing Server: %v", app.Listen(serverAddr, fiber.ListenConfig{EnablePrefork: true}))
}
