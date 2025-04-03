package main

import (
	"backend/db"
	"backend/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	g := gin.Default()

	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	g.Use(cors.New(cors.Config{
		AllowAllOrigins: true,           // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"*"}, // Allow all headers
		ExposeHeaders:     []string{"Content-Type", "X-Requested-With"}, // Expose specific headers if needed
		AllowCredentials: false,         // Optional, set to true if you need credentials (cookies, auth headers)
	}))

	routes.SetupAllRoutes(g)

	db.InitDB()	

	g.Run(":8080")
}