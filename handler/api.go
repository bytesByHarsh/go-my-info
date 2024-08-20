package handler

import (
	"log"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/gofiber/fiber/v3"
)

type apiConfig struct {
	DB *database.Queries
}

var apiCfg apiConfig

func UpdateDB(db *database.Queries) {
	if db == nil {
		log.Fatal("Empty DB pointer Received")
	}
	apiCfg.DB = db
}

// Hello handle api status
func Hello(c fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm not ok!", "data": nil}, "")
}
