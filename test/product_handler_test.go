package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"test-golang-muhamad-isro/database"
	"test-golang-muhamad-isro/entity/domain"
	"test-golang-muhamad-isro/helper"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

//	func truncateCategory(db *gorm.DB) {
//		db.Exec("TRUNCATE category")
//	}
func TestFindAllProducts(t *testing.T) {
	app := fiber.New()

	db := database.OpenConnection()

	defer database.CloseConnection(db)

	// db.Create(&domain.Product{
	// 	ID:          uuid.New(),
	// 	Name:        "Test Product",
	// 	Description: "Test Description",
	// 	CategoryID:  uuid.New(),
	// })

	req := httptest.NewRequest("GET", "/api/products/", nil)
	resp, error := app.Test(req)

	helper.PanicIfError(error)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var products []domain.Product
	json.NewDecoder(resp.Body).Decode(&products)

	fmt.Println(products)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
