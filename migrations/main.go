package main

import (
	"app/database"
	"app/internal/domain/category"
	"app/internal/domain/product"
)

func main() {
	db := database.Connect()

	db.AutoMigrate(&category.Category{}, &product.Product{})
}
