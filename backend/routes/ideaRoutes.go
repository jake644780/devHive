package routes

import (
	"github.com/gin-gonic/gin"
	"backend/controllers"
)

func SetupIdeaRoutes(r *gin.Engine){
	r.GET("/ideas", func(c *gin.Context) {
		controllers.IdeaController(c, "all")
	})
	r.GET("/idea/:id", func(c *gin.Context) {
		controllers.IdeaController(c, "get")
	})
	r.POST("/ideas/new", func(c *gin.Context) {
		controllers.IdeaController(c, "create")
	})
	r.DELETE("/idea/:id", func(c *gin.Context) {
		controllers.IdeaController(c, "delete")
	})

}