package dtos

type CreatMenuDTO struct {
	Name string `json:"name" binding:"required"`
}

type MenuDTO struct {
	Name string `json:"name"`
}
