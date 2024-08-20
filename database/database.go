package db

import (
	"database/sql"
	"log"

	"github.com/bytesByHarsh/go-my-info/config"
	"github.com/bytesByHarsh/go-my-info/handler"
	"github.com/bytesByHarsh/go-my-info/internal/database"

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

	return nil
}
