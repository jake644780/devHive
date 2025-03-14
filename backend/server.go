package main

import (
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	routes.SetupRoutes(g);


	g.Run(":8081");
}