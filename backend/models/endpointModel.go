package models

import "gorm.io/gorm"

type Endpoint struct {
	gorm.Model
	Route string `gorm:"unique"`
	Name  string
}
