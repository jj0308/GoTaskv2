package main

import (
	"github.com/jj0308/GoTaskv2/config"
	"github.com/jj0308/GoTaskv2/storage"
	"github.com/jj0308/GoTaskv2/routes"
	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	cfg := config.LoadConfig()

	db := storage.SetupDatabase(cfg)
	defer db.Close()

	router := router.SetupRouter(db)
	router.Run(cfg.Server.Address)
}
