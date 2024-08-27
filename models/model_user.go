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

type CreateUserByAdminReq struct {
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Password    string `json:"password" validate:"required"`
	IsSuperUser bool   `json:"is_superuser"`
	IsActive    bool   `json:"is_active" validate:"required"`
}

type UpdateUserReq struct {
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Name       string `json:"name" validate:"required"`
	PhoneNum   string `json:"phone_num" validate:"required"`
	ProfileImg string `json:"profile_img" validate:"required"`
}

type UpdatePasswordReq struct {
	Password string `json:"password" validate:"required"`
}

type GetUserListReq struct {
	Page         int `json:"page"`
	ItemsPerPage int `json:"items_per_page"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	PhoneNum    string    `json:"phone_number"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	ProfileImg  string    `json:"profile_img"`
	Role        int32     `json:"role"`
	IsSuperUser bool      `json:"is_superuser"`
	IsActive    bool      `json:"is_active"`
}

func ConvUserToUser(dbUser database.User) User {
	return User{
		ID:          dbUser.ID,
		CreatedAt:   dbUser.CreatedAt,
		UpdatedAt:   dbUser.UpdatedAt,
		Name:        dbUser.Name,
		PhoneNum:    dbUser.PhoneNum,
		Email:       dbUser.Email,
		Username:    dbUser.Username,
		ProfileImg:  dbUser.ProfileImg,
		Role:        dbUser.Role,
		IsSuperUser: dbUser.Role == 100,
		IsActive:    dbUser.IsActive,
	}
}

func CreateUserListResp(dbUserList []database.User) []User {
	userList := []User{}
	for _, dbUser := range dbUserList {
		userList = append(userList, ConvUserToUser(dbUser))
	}
	return userList
}
