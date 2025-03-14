package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupAllRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "helo word",
		});
	});
	SetupIdeaRoutes(r)
	log.Println("routes set up")
}