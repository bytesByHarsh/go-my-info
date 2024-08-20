package main

import (
	"fmt"
	"log"

	"github.com/bytesByHarsh/go-my-info/config"
	db "github.com/bytesByHarsh/go-my-info/database"
	"github.com/bytesByHarsh/go-my-info/router"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type structValidator struct {
	validate *validator.Validate
}

// Validator needs to implement the Validate method
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive:   false,
		StrictRouting:   false,
		ServerHeader:    "Fiber",
		AppName:         "My Info",
		StructValidator: &structValidator{validate: validator.New()},
	})

	config.ReadEnvFile(".env")

	// log.Printf("Starting Server at: %v:%v", config.Cfg.SERVER_LINK, config.Cfg.SERVER_PORT)

	serverAddr := fmt.Sprintf("%v:%v", config.Cfg.SERVER_LINK, config.Cfg.SERVER_PORT)

	db.ConnectDb()

	router.SetupRoutes(app)
	log.Fatalf("Closing Server: %v", app.Listen(serverAddr, fiber.ListenConfig{EnablePrefork: true}))
}
