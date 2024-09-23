package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/bytesByHarsh/go-my-info/config"
	"github.com/bytesByHarsh/go-my-info/handler"
	"github.com/bytesByHarsh/go-my-info/internal/database"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

var DB *database.Queries

func ConnectDb() error {
	dbConn, err := sql.Open("postgres", config.Cfg.DB_URL)
	if err != nil {
		log.Fatal("Can't connect to database:", err)
	}
	DB = database.New(dbConn)

	handler.UpdateDB(DB)

	InitDb()

	return nil
}

func InitDb() {
	adminDetails := database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		DeletedAt:      sql.NullTime{},
		IsDeleted:      false,
		Name:           "Admin User",
		Email:          "admin@admin.com",
		PhoneNum:       "+919879879",
		ProfileImg:     "",
		Username:       "admin",
		Role:           handler.UserRole_Admin,
		HashedPassword: handler.HashPassword("123"),
	}
	_, err := DB.GetUserByUsername(context.Background(), adminDetails.Username)
	if err == nil {
		// Admin user already created
		return
	}

	_, err = DB.CreateUser(context.Background(), adminDetails)
	if err != nil {
		log.Fatalln("Not Able to create admin credentials")
	}

}
