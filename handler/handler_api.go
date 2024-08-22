package handler

import (
	"log"
	"net/http"

	"github.com/bytesByHarsh/go-my-info/config"
	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/bytesByHarsh/go-my-info/models"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
)

type apiConfig struct {
	DB        *database.Queries
	AuthToken *jwtauth.JWTAuth
}

var apiCfg apiConfig

func UpdateDB(db *database.Queries) {
	if db == nil {
		log.Fatal("Empty DB pointer Received")
	}
	apiCfg.DB = db
}

func Init() {
	models.Validate = validator.New()

	apiCfg.AuthToken = jwtauth.New("HS256", []byte(config.Cfg.JWT_SECRET_KEY), nil)

	claims := map[string]interface{}{
		"id": "123456789",
		// "username": "admin",
		// "role":     10,
	}
	// For confirming that tokens are generating
	_, _, err := apiCfg.AuthToken.Encode(claims)
	if err != nil {
		log.Fatalf("Token not generated: %v", err)
	}
	// log.Printf("Dummy Token: %v", tokenString)
}

// Hello handle api status
func Hello(w http.ResponseWriter, r *http.Request) {
	resp := models.JSONResp{
		Status:  "success",
		Message: "Hello i'm not ok!",
		Data:    nil,
	}
	responseWithJson(w, 201, resp)
}
