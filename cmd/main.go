package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bytesByHarsh/go-my-info/config"
	db "github.com/bytesByHarsh/go-my-info/database"
	"github.com/bytesByHarsh/go-my-info/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	config.ReadEnvFile(".env")

	// log.Printf("Starting Server at: %v:%v", config.Cfg.SERVER_LINK, config.Cfg.SERVER_PORT)

	serverAddr := fmt.Sprintf("%v:%v", config.Cfg.SERVER_LINK, config.Cfg.SERVER_PORT)

	db.ConnectDb()

	app := chi.NewRouter()
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	app.Use(middleware.RequestID)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	// app.Use(middleware.StripSlashes)
	app.Use(middleware.Heartbeat("/ping"))

	router.SetupRoutes(app)
	log.Printf("Server Starting on Address: %v", serverAddr)
	err := http.ListenAndServe(serverAddr, app)
	if err != nil {
		log.Fatal(err)
	}
}
