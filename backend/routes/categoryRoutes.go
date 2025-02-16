package routes

import (
	"gomenu/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.POST("/", controllers.CreateCategory)
		categoryRoutes.GET("/", controllers.GetAllCategories)
	}
}
