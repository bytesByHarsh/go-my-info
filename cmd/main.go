package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bytesByHarsh/go-my-info/api"
	"github.com/bytesByHarsh/go-my-info/config"
	db "github.com/bytesByHarsh/go-my-info/database"
	"github.com/bytesByHarsh/go-my-info/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @contact.name   Harsh Mittal
// @contact.url
// @contact.email  harshmittal2210@gmail.com

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

	// Swagger
	api.SwaggerInfo.Title = "My Information Server"
	api.SwaggerInfo.Version = "0.0.1"
	api.SwaggerInfo.BasePath = "/v1"
	api.SwaggerInfo.Schemes = []string{"http", "https"}
	api.SwaggerInfo.Description = ` Simple Backend API to manage your Accounts, Cards etc.

	Ping: To test if server is UP
	Authentication: Auth related API
	Banks: Banks API
	Accounts: Bank Account API
	Cards: Debit/Credit Card API

`

	app.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition"
	))

	router.SetupRoutes(app)
	log.Printf("Server Starting on Address: %v", serverAddr)
	err := http.ListenAndServe(serverAddr, app)
	if err != nil {
		log.Fatal(err)
	}
}
