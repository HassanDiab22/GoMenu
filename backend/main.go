package main

import (
	"gomenu/controllers"
	"gomenu/initializers"
	"gomenu/middleware"
	"gomenu/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()

	r.POST("/signin", controllers.Signin)
	r.POST("/register", controllers.Register)

	authRoutes := r.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())

	{
		routes.CategoryRoutes(authRoutes) // Load category routes
		routes.MenuRoutes(authRoutes)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
