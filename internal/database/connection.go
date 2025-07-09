package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/lasamarndi1994/gov2/internal/config"
)

// DBClient represents our database connection pool
var DB *sql.DB

// InitDB initializes the database connection pool
func InitDB(cfg *config.Config) {
	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBHost, cfg.DBPort, cfg.DBName)

	fmt.Println("Attempting to connect to database:", cfg.DBName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Configure connection pool
	DB.SetMaxOpenConns(100)                // Max 100 open connections
	DB.SetMaxIdleConns(25)                 // Max 25 idle connections
	DB.SetConnMaxLifetime(5 * time.Minute) // Connections will be reused for at most 5 minutes
	DB.SetConnMaxIdleTime(1 * time.Minute) // Connections idle for 1 minute will be closed

	// Ping the database to verify the connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to MySQL database!")
}

// CloseDB closes the database connection pool
func CloseDB() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			fmt.Println("Database connection closed.")
		}
	}
}

// GetDB returns the initialized database connection pool
func GetDB() *sql.DB {
	return DB
}
