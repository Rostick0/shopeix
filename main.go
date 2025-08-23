package main

import (
	"app/database"
	"app/internal/domain/category"
	CategoryRepository "app/internal/domain/category/repository"
	CategoryService "app/internal/service/category"
	transportHttp "app/internal/transport/http"
	CategoryTransport "app/internal/transport/http/category"
	"fmt"
	"net/http"
	// "app/internal/domain/category"
)

func main() {
	db := database.Connect()

	// categoryRepo := category.NewService()()
	db.AutoMigrate(&category.Category{})

	categoryRepository := CategoryRepository.NewCategoryGormRepo(db)
	categoryService := CategoryService.NewService(categoryRepository)
	categoryHandler := CategoryTransport.NewCategoryHandler(categoryService)

	r := transportHttp.NewRouter(categoryHandler)
	// r.ListenAndServe()
	fmt.Print("Start server: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
