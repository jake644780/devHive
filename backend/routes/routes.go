package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "helo word",
		});
	});
	SetupIdeaRoutes(r)
}