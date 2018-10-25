package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zuhrulumam/crud-gorm/models"
)

var db *gorm.DB
var err error
var id uint

// type Product struct {
// 	gorm.Model
// 	Category      Category `gorm:"foreignkey:CategoryRefer"`
// 	Code          string
// 	Price         uint
// 	CategoryRefer uint
// }

// type Category struct {
// 	gorm.Model
// 	Name     string    `gorm:"unique"`
// 	Products []Product `gorm:"foreignkey:CategoryRefer"`
// }

func main() {
	db, err = gorm.Open("mysql", "ritx:ritx@tcp(localhost:3306)/ritx?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	fmt.Println("done connecting")

	defer db.Close()

	db.DropTableIfExists(models.Product{}, models.Category{})

	db.CreateTable(models.Product{}, models.Category{})

	createProduct()
	// findProducts()
	// findProduct()
	// updateProductCategory()
	findProductWithAssociation()
	// findCategory()
	// updateProduct()
	// deleteProduct()
}

// add before create and others

func findProductWithAssociation() {
	var product []models.Product

	// it'll always return all products but only preload the condition
	if err := db.Debug().Model(models.Product{}).Preload("Category", "name = ?", "Test Category").Find(&product).Error; err != nil {
		panic(err)
	}

	fmt.Println("Product Association ", product)

	for _, value := range product {
		fmt.Println(value.Category.Name)
	}

	// it's only work for retrieved record
	// var category Category
	// db.Model(&product[0]).Related(&category)

	// fmt.Println("Category Association ", category)

}

func updateProductCategory() {
	var product models.Product

	if err := db.Model(models.Product{}).Preload("Category").First(&product).Error; err != nil {
		panic(err)
	}

	product.Category.Name = "Test Updated Category via Product2"

	db.Model(models.Product{}).Save(&product)

}

// find one by id with pages and sort
func findProduct() {
	var product models.Product

	if err := db.Model(models.Product{}).Preload("Category").Where("id = ?", id).Limit(1).Offset(0).Order("price desc").Find(&product).Error; err != nil {
		panic(err)
	}

	fmt.Println("Product ", product)
	fmt.Println("Category ", product.Category)
}

// func createCategory() {
// 	category := Category{name: "Test Category"}

// 	db.Model(models.Category{}).Create(&category)
// }

func findCategory() {
	var categories []models.Category
	if err := db.Model(models.Category{}).Find(&categories).Error; err != nil {
		panic("something wrong")
	}

	fmt.Println("List Categories ", categories)
}

// create
func createProduct() {
	product := models.Product{Code: "testProduct", Price: 12000, Category: models.Category{Name: "Test Category 2"}}

	db.Model(models.Product{}).Create(&product)
	db.Model(models.Product{}).Save(&product)

	// // fmt.Println("id ", product.ID)
	// id = product.ID

	// product2 := Product{Code: "testProduct2", Price: 12000, Category: Category{Name: "Test Category"}}

	// db.Model(models.Product{}).Create(&product2)
	// db.Model(models.Product{}).Save(&product2)

	// product3 := Product{Code: "testProduct3", Price: 12000, Category: Category{Name: "Test Category"}}

	// db.Model(&product3).Association("Category").Append(Category{Name: "Test Category"})
	// db.Model(models.Product{}).Create(&product3)

	// instead of using product we use categories to create products

	products := models.Category{
		Name: "Test Category",
		Products: []models.Product{
			{
				Code:  "Product 1",
				Price: 120000,
			},
			{
				Code:  "Product 2",
				Price: 15000,
			},
		},
	}

	db.Debug().Model(models.Category{}).Create(&products)
	db.Model(models.Category{}).Save(&products)

	fmt.Println("Products ", products)
}

// find
func findProducts() {
	var products []models.Product
	if err := db.Debug().Model(models.Product{}).Preload("Category").Find(&products).Error; err != nil {
		panic("something wrong")
	}

	fmt.Println("List Products ", products)
	fmt.Println("Products ", products[0].Category.Name)

	// var category Category
	var product models.Product

	id = products[0].ID

	db.Preload("Category").First(&product)

	// db.Debug().Model(&product).Related(&category)
	// db.Debug().Preload("Category").
	fmt.Println("Category ", product.Category.Name)
	// // find with pages
}

// update
func updateProduct() {
	var product models.Product

	if err := db.Model(models.Product{}).Where("id = ?", id).First(&product).Error; err != nil {
		panic(err)
	}

	product = models.Product{Code: "Updated", Price: 15000}

	db.Model(models.Product{}).Save(&product)

	fmt.Println("Updated Product ", product)

}

// delete
func deleteProduct() {
	var deletedProduct models.Product

	product := db.Model(models.Product{}).Where("id = ?", id).Delete(&deletedProduct)

	fmt.Println("Deleted Product ", product)

}
