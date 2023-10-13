package database

import (
	"api/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB(config config.Config) (*sql.DB, error) {

	// Construct the connection string
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		config.DBUser, config.DBPassword, config.DBName, config.SSLMode, config.DBHost, config.DBPort)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	// Try to ping the database to check the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	return db, nil
}
