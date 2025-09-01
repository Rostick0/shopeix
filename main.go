package main

import (
	"app/database"
	CategoryRepository "app/internal/domain/category/repository"
	ProductRepository "app/internal/domain/product/repository"
	CategoryService "app/internal/service/category"
	ProductService "app/internal/service/product"
	TransportHttp "app/internal/transport/http"
	CategoryTransport "app/internal/transport/http/category"
	ProductTransport "app/internal/transport/http/product"
	"fmt"
	"net/http"
	// "github.com/go-playground/validator/v10"
)

func main() {
	db := database.Connect()
	// validate := validator.New()

	categoryRepository := CategoryRepository.NewGormRepo(db)
	categoryService := CategoryService.NewService(categoryRepository)
	categoryHandler := CategoryTransport.NewHandler(categoryService)

	productRepository := ProductRepository.NewGormRepo(db)
	productService := ProductService.NewService(productRepository)
	productHandler := ProductTransport.NewHandler(productService)

	r := TransportHttp.NewRouter(categoryHandler, productHandler)
	// r.ListenAndServe()
	fmt.Print("Start server: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
