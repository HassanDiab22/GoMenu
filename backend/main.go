package main

import (
	"fmt"
	"gomenu/controllers"
	"gomenu/initializers"
	"gomenu/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	fmt.Println("in maoin")
	r.POST("/signin", controllers.Signin)
	r.POST("/register", controllers.Register)
	r.POST("/category", controllers.CreateCategory)
	r.GET("/category", middleware.AuthMiddleware(), controllers.GetAllCategories)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
