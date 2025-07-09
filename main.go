package main

import (
	"github.com/lasamarndi1994/gov2/api/routes"
	"github.com/lasamarndi1994/gov2/internal/config"
	"github.com/lasamarndi1994/gov2/internal/database"
)

func main() {
	cfg := config.LoadConfig() //Load configuration from .env
	database.InitDB(cfg)       // init database
	defer database.CloseDB()   // close database connection

	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + cfg.AppPort)
}
