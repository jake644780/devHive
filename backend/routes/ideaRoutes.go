package routes

import (
	"backend/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupIdeaRoutes(r *gin.Engine){
	r.GET("/ideas", controllers.GetAllIdeas)
	r.GET("/idea/:id", controllers.GetIdea)
	r.POST("/idea/new", controllers.CreateNewIdea)
	r.DELETE("/idea/:id", controllers.DeleteIdea)
	log.Println("\tideaRoutes")
}