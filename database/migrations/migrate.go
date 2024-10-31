// database/migrations/migrate.go
package main

import (
	"fmt"
	"log"
	"product-store/app/models"
	"product-store/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Database connection
	dsn := config.DatabaseDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Run migrations for User and other models
	if err := db.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		log.Fatal("failed to migrate tables:", err)
	}

	fmt.Println("Migration completed successfully.")
}
