package models

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

type JSONResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func VerifyJson(data interface{}, r *http.Request) error {

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}

	// Validate the struct
	if err := Validate.Struct(data); err != nil {
		return err
	}

	return nil
}
