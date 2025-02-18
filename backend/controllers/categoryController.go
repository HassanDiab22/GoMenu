package controllers

import (
	"gomenu/dtos"
	"gomenu/initializers"
	"gomenu/models"
	"gomenu/utils"
	"gomenu/validations"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCategories(c *gin.Context) {
	var categories []models.Category
	var categoriesDTO []dtos.CategoryDTO
	// var categoryDTOs []dtos.CategoryDTO

	initializers.DB.Find(&categories)

	if err := utils.AutoMap(categories, &categoriesDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"categories": categoriesDTO,
	})
}

func CreateCategory(c *gin.Context) {
	var body dtos.CreateCategoryDTO

	_, exists := c.Get("validatedData")
	if !exists {
		c.JSON(500, gin.H{"error": "Validation middleware failed"})
		return
	}

	if err := c.ShouldBind(&body); err != nil {
		utils.Validate(err, validations.CategoryCustomValidationMessages, c)
		return
	}

	category := models.Category{Name: body.Name, Index: body.Index}
	initializers.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}
