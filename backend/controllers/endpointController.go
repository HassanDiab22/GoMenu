package controllers

import (
	"gomenu/dtos"
	"gomenu/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EndpointCategory(c *gin.Context) {
	var body dtos.CreateEndpointDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(500, gin.H{"error": "Validation middleware failed"})
		return
	}

	endpoint, err := services.CreateEndpoint(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create endpoint"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Endpoint created successfully",
		"endpoint": endpoint,
	})
}
