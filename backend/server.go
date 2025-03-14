package main

import (
	"backend/db"
	"backend/routes"
	"log"

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
	routes.SetupRoutes(g)

	db, err := db.InitDB();
	if err != nil{
		log.Fatal(err)
		return
	}
	log.Printf("%v", db)
	

	g.Run(":8081")
}