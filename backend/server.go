package main

import (
	"backend/routes"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	g := gin.Default()

	// Load environment variables from the .env file
	err := godotenv.Load();
	if err != nil {
		log.Fatal("Error loading .env file");
	}
	routes.SetupRoutes(g);



	g.Run(":8081");
}