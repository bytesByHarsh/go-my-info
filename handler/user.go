package handler

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"github.com/bytesByHarsh/go-my-info/config"
	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/bytesByHarsh/go-my-info/models"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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

	start := time.Now()
	hash := hashPassword(user.Password)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusBadRequest, "Password Input issue")
	// }
	end := time.Now()
	fmt.Printf("myFunction took %v to complete.\n", end.Sub(start))

	start = time.Now()
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
	end = time.Now()
	fmt.Printf("myFunction took %v to complete.\n", end.Sub(start))

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("Err: %v", err))
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Created user",
		"data":    models.ConvUserToUser(dbUser),
	})
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
