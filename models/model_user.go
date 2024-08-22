package models

import (
	"time"

	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/google/uuid"
)

type CreateUserReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	PhoneNum   string    `json:"phone_number"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	ProfileImg string    `json:"profile_img"`
	Role       int32     `json:"role"`
}

func ConvUserToUser(dbUser database.User) User {
	return User{
		ID:         dbUser.ID,
		CreatedAt:  dbUser.CreatedAt,
		UpdatedAt:  dbUser.UpdatedAt,
		Name:       dbUser.Name,
		PhoneNum:   dbUser.PhoneNum,
		Email:      dbUser.Email,
		Username:   dbUser.Username,
		ProfileImg: dbUser.ProfileImg,
		Role:       dbUser.Role,
	}
}
