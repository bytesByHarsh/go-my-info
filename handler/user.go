package handler

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/bytesByHarsh/go-my-info/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	user := new(NewUser)

	if err := c.Bind().Body(user); err != nil {
		log.Printf("Parse Error: %v", err)
		return fiber.NewError(fiber.StatusBadRequest, "wrong values sent")
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Password Input issue")
	}

	// log.Printf("User Details: %v", user)
	dbUser, err := apiCfg.DB.CreateUser(c.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		DeletedAt:      sql.NullTime{},
		IsDeleted:      false,
		Name:           user.Name,
		Email:          user.Email,
		Username:       user.Username,
		PhoneNum:       "",
		ProfileImg:     "",
		Role:           10,
		HashedPassword: hash,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("Err: %v", err))
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created user",
		"data":    models.ConvUserToUser(dbUser),
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
