// main.go
package main

import (
	"log"
	"product-store/config"
	"product-store/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	dsn := config.DatabaseDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	return db
}

func main() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	db := initDB()

	// Set up routes
	r := routes.SetupRouter(db)

	// Run server
	r.Run(":8082")
}
