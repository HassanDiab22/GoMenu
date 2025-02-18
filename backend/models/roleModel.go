package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"unique"`
	Endpoints []Endpoint `gorm:"many2many:role_endpoints;"`
}
