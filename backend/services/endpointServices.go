package services

import (
	"gomenu/dtos"
	"gomenu/initializers"
	"gomenu/models"
)

func CreateEndpoint(data dtos.CreateEndpointDTO) (models.Endpoint, error) {
	endpoint := models.Endpoint{Name: data.Name, Route: data.Route}
	if err := initializers.DB.Create(&endpoint).Error; err != nil {
		return models.Endpoint{}, err
	}
	return endpoint, nil
}
