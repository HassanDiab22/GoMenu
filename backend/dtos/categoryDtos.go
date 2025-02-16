package dtos

type CreateCategoryDTO struct {
	Name  string `json:"name" binding:"required"`
	Index int    `json:"index" binding:"required,gte=0"`
}

type CategoryDTO struct {
	ID    int
	Name  string
	Index int
}
