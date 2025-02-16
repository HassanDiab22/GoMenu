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
	// var categoryDTOs []dtos.CategoryDTO

	initializers.DB.Find(&categories)

	// for _, category := range categories {
	// 	categoryDTOs = append(categoryDTOs, dtos.CategoryDTO{
	// 		Name: category.Name,
	// 		Index: category.Index,
	// 		ID: CATE,
	// 	})
	// }

	c.JSON(200, gin.H{
		"categories": categories,
	})
}

func CreateCategory(c *gin.Context) {
	var body dtos.CreateCategoryDTO

	// ✅ Use ShouldBindJSON to validate and return errors
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.Validate(err, validations.CategoryCustomValidationMessages, c)
		return
	}

	category := models.Category{Name: body.Name, Index: body.Index}
	initializers.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{
		"category": category,
	})
}
