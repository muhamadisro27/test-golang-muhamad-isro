package main

import (
	"fmt"
	"test-golang-muhamad-isro/database"
	handler "test-golang-muhamad-isro/handler/product"
	repository "test-golang-muhamad-isro/repository/product"
	"test-golang-muhamad-isro/router"
	service "test-golang-muhamad-isro/service/product"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.OpenConnection()
	defer database.CloseConnection(db)

	app := fiber.New()
	port := 4000
	validate := validator.New()


	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productHandler := handler.NewProductHandler(productService)

	router.ProductRouter(app, productHandler)


	fmt.Printf("Server listening on port %d...\n", port)
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
