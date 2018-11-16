package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	// Category      Category `gorm:"foreignkey:CategoryRefer"`
	Code  string
	Price uint
	// CategoryRefer uint
	Quantity uint
	// TotalPrice    uint
}

// it's not affected association
// func (product *Product) BeforeSave(tx *gorm.DB) (err error) {
// 	totalPrice := product.Quantity * product.Price
// 	tx.Model(&Product{}).Update("total_price", totalPrice)

// 	return
// }
