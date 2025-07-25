package main

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/lasamarndi1994/gov2/api/routes"
	"github.com/lasamarndi1994/gov2/internal/config"
	"github.com/lasamarndi1994/gov2/internal/database"
)

var validate *validator.Validate

func main() {
	cfg := config.LoadConfig() //Load configuration from .env
	database.InitDB(cfg)       // init database
	defer database.CloseDB()   // close database connection
	formValidation()           // add form validation
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + cfg.AppPort)
}

func formValidation() {
	validate = validator.New()

	// Register function to get field name from `json` tag
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}
