package test

import (
	"net/http"
	"net/http/httptest"
	"test-golang-muhamad-isro/database"
	handler "test-golang-muhamad-isro/handler/product"
	repository "test-golang-muhamad-isro/repository/product"
	"test-golang-muhamad-isro/router"
	service "test-golang-muhamad-isro/service/product"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupTestApp() (*fiber.App, *gorm.DB) {
	// Initialize a new Fiber app
	app := fiber.New()

	db := database.OpenConnection()

	validate := validator.New()

	// Define routes
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productHandler := handler.NewProductHandler(productService)

	router.ProductRouter(app, productHandler)

	return app, db
}
func TestFindAllProducts(t *testing.T) {

	app, db := setupTestApp()

	defer database.CloseConnection(db)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
