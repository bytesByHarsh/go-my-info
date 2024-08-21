package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	SERVER_PORT string
	SERVER_LINK string
	DB_URL      string

	SECRET_KEY string
}

var Cfg EnvConfig

func ReadEnvFile(envPath string) {
	if envPath == "" {
		envPath = ".env"
	}
	// load .env file
	err := godotenv.Load(envPath)
	if err != nil {
		fmt.Print("Error loading .env file")
	}

	Cfg.SERVER_LINK = os.Getenv("SERVER_LINK")
	Cfg.SERVER_PORT = os.Getenv("SERVER_PORT")
	Cfg.DB_URL = os.Getenv("DB_URL")
	Cfg.SECRET_KEY = os.Getenv("SECRET_KEY")

	if Cfg.SERVER_LINK == "" {
		Cfg.SERVER_LINK = "localhost"
	}

	if Cfg.SERVER_PORT == "" {
		Cfg.SERVER_PORT = "8005"
	}

	if Cfg.DB_URL == "" {
		log.Fatal("DB URL is not Mentioned")
	}
}

func Config(key string) string {
	return os.Getenv(key)
}
