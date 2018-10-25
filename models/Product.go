package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Category      Category `gorm:"foreignkey:CategoryRefer"`
	Code          string
	Price         uint
	CategoryRefer uint
}
