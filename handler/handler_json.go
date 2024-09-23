package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bytesByHarsh/go-my-info/models"
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON Response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %v", msg)
	}

	responseWithJson(w, code, models.JSONerrResponse{
		Error: msg,
	})
}
