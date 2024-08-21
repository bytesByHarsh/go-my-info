package handler

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bytesByHarsh/go-my-info/config"
	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/bytesByHarsh/go-my-info/models"
	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	params := parameters{}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	// Validate the struct
	if err := validate.Struct(params); err != nil {
		responseWithError(w, http.StatusBadRequest,
			fmt.Sprintf("Error parsing JSON: %v", err),
		)
		return
	}

	hash := hashPassword(params.Password)

	dbUser, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		DeletedAt:      sql.NullTime{},
		IsDeleted:      false,
		Name:           params.Name,
		Email:          params.Email,
		Username:       params.Username,
		PhoneNum:       "",
		ProfileImg:     "",
		Role:           10,
		HashedPassword: hash,
	})

	if err != nil {
		responseWithError(w, 400,
			fmt.Sprintf("couldn't create user: %v", err),
		)
		return
	}

	resp := models.JSONResp{
		Status:  "success",
		Message: "User Created",
		Data:    models.ConvUserToUser(dbUser),
	}
	responseWithJson(w, 201, resp)
}

func hashPassword(password string) string {
	// bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// return string(bytes), err
	// Concatenate the secret and password
	combined := config.Cfg.SECRET_KEY + password
	hash := sha256.New()
	hash.Write([]byte(combined))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}
