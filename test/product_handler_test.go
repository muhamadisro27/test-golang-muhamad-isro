package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"test-golang-muhamad-isro/database"
	"test-golang-muhamad-isro/entity/domain"
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

func truncateCategory(db *gorm.DB) {
	db.Exec("TRUNCATE products")
	db.Exec("TRUNCATE product_categories")
}

func TestInsertNewProduct(t *testing.T) {

	app, db := setupTestApp()
	truncateCategory(db)
	defer database.CloseConnection(db)

	category := domain.ProductCategory{
		Name: "Category 1",
	}

	err := db.Create(&category).Error
	assert.Nil(t, err)

	var getCategory domain.ProductCategory
	err = db.First(&getCategory, "id = ?", category.ID).Error
	assert.Nil(t, err)

	requestBody := map[string]interface{}{
		"name":        "test product 2",
		"description": "test",
		"category_id": category.ID,
	}

	jsonBody, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/products/", bytes.NewBuffer(jsonBody))

	res, err := app.Test(req)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(res.Body)

	fmt.Println(string(bytes))

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
}

func TestFindAllProducts(t *testing.T) {

	app, db := setupTestApp()

	defer database.CloseConnection(db)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/products/", nil)
	res, err := app.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
}
