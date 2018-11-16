package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"unique"`
	// Products []Product `gorm:"foreignkey:CategoryRefer"`
}

// func (category *Category) BeforeCreate(tx *gorm.DB) (err error) {
// 	fmt.Println("Before create Category")
// 	for _, product := range category.Products {
// 		product.TotalPrice = product.Quantity * product.Price
// 	}
// 	return
// }
