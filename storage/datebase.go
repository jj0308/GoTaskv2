package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jj0308/GoTaskv2/config"
	_ "github.com/denisenkom/go-mssqldb"
)

func SetupDatabase(cfg *config.Config) *sql.DB {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;Database=%s",
		cfg.Database.Server, cfg.Database.User, cfg.Database.Password, cfg.Database.Port, cfg.Database.Database)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	return db
}

