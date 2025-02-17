package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name string
	Url  string
}
