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

func CreatMenu(c *gin.Context) {
	var requestData dtos.CreatMenuDTO

	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.Validate(err, validations.CreatMenuValidationMessages, c)
		return
	}

	menu := models.Menu{Name: requestData.Name}
	if err := initializers.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"menu": menu,
	})

}

func GetAllMenus(c *gin.Context) {
	var menus []models.Menu

	total, err := utils.Paginate(initializers.DB, &menus, c.Query("page"), c.Query("page_size"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menus"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":     total,
		"page":      c.Query("page"),
		"page_size": c.Query("page_size"),
		"menus":     menus,
	})
}
