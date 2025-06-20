package db

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/abhinavgupta21/go-ci-cd-project/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB function opens DB connection using GORM and performs AutoMigrate.
func ConnectDB() (*gorm.DB, error) {
	connectionStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPassword,
		config.DBName, config.DBPort, config.DBSSLMode,
	)

	gormDB, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Connected to database")
	return gormDB, nil
}
