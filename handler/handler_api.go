package handler

import (
	"log"
	"net/http"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/bytesByHarsh/go-my-info/models"
	"github.com/go-playground/validator/v10"
)

type apiConfig struct {
	DB *database.Queries
}

var apiCfg apiConfig
var validate *validator.Validate

func UpdateDB(db *database.Queries) {
	if db == nil {
		log.Fatal("Empty DB pointer Received")
	}
	apiCfg.DB = db
}

func Init() {
	validate = validator.New()
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
