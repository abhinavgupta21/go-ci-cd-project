package db

import (
	"fmt"
	"log"

	"github.com/abhinavgupta21/go-ci-cd-project/models"

	"gorm.io/gorm"
)

func UpdateTables(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	log.Println("Database migrated successfully")
	return nil
}
