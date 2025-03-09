package main

import (
	"fmt"
	"go-url-shortener/config"
	"go-url-shortener/database" // Import generated Swagger docs
	"go-url-shortener/modules/url"
	"go-url-shortener/routes"
	"log"
	"os"
)

// @title Golang URL Shortener API
// @version 1.0
// @description A simple URL shortener service built with Gin and PostgreSQL
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api/v1
func main() {
	config.LoadEnv()
	database.Connect()
	database.ConnectRedis()

	// Perform database migration
	err := database.DB.AutoMigrate(&url.ShortenedURL{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("Database migration completed successfully!")

	router := routes.SetupRouter()

	port := os.Getenv("SERVER_PORT")
	fmt.Println("Server is running on port:", port)
	log.Fatal(router.Run(":" + port))
}
