package dtos

type CreateEndpointDTO struct {
	Name  string `json:"name" binding:"required"`
	Route string `json:"route" binding:"required"`
}
