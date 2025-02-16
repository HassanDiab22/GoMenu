package main

import (
	"fmt"
	"gomenu/controllers"
	"gomenu/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	fmt.Println("in maoin")
	r.POST("/category", controllers.CreateCategory)
	r.GET("/category", controllers.GetAllCategories)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
