package routes

import (
	"gomenu/controllers"
	"gomenu/dtos"
	"gomenu/middleware"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup) {
	categoryRoutes := r.Group("/category")
	{
		categoryRoutes.POST("/", middleware.ValidationMiddleware(dtos.CreateCategoryDTO{}), controllers.CreateCategory)
		categoryRoutes.GET("/", controllers.GetAllCategories)
	}
}
