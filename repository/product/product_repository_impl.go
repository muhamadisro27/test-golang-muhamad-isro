package repository

import (
	"test-golang-muhamad-isro/entity/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Insert(DB *gorm.DB, product domain.Product) domain.Product {
	tx := DB.Begin()
	defer tx.Rollback()

	query := tx

	err := query.Create(&product).Error

	if err == nil {
		tx.Commit()
	}

	return product
}

func (repository *ProductRepositoryImpl) FindAll(ctx *fiber.Ctx, DB *gorm.DB) []domain.Product {

	products := []domain.Product{}

	keyword := ctx.Query("keyword")
	category := ctx.Query("category")

	tx := DB.Begin()
	defer tx.Rollback()

	query := tx.Preload("Category")

	if keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if category != "" {
		query = query.Where("category_id = ?", category)
	}

	err := query.Find(&products).Error

	if err == nil {
		tx.Commit()
	}

	return products
}
