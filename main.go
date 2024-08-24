package main

import (
	"fmt"
	"test-golang-muhamad-isro/config"
	"test-golang-muhamad-isro/database"
	handler "test-golang-muhamad-isro/handler/product"
	repository "test-golang-muhamad-isro/repository/product"
	"test-golang-muhamad-isro/router"
	service "test-golang-muhamad-isro/service/product"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.InitConfig()

	DATABASE_URL := config.GetString("DATABASE_URL")

	db := database.OpenConnection(DATABASE_URL)
	defer database.CloseConnection(db)

	app := fiber.New()
	port := config.GetString("PORT")
	validate := validator.New()

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productHandler := handler.NewProductHandler(productService)

	router.ProductRouter(app, productHandler)

	fmt.Printf("Server listening on port %s...\n", port)
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}
