package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	Code string
	Name string
	Category string
	Quantity string
}