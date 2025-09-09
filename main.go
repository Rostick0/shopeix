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

	"github.com/joho/godotenv"
	// "github.com/go-playground/validator/v10
)

var serverPort = "8080"

func main() {
	godotenv.Load()
	db := database.Connect()

	categoryRepository := CategoryRepository.NewGormRepo(db)
	categoryService := CategoryService.NewService(categoryRepository)
	categoryHandler := CategoryTransport.NewHandler(categoryService)

	productRepository := ProductRepository.NewGormRepo(db)
	productService := ProductService.NewService(productRepository)
	productHandler := ProductTransport.NewHandler(productService)

	r := TransportHttp.NewRouter(categoryHandler, productHandler)
	fmt.Print("Start server: http://localhost:" + serverPort)
	http.ListenAndServe(":"+serverPort, r)
}
