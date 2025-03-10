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
