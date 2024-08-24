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

	query := tx

	err := query.Create(&product).Error

	if err != nil {
		tx.Rollback()
		panic(err)
	}

	err = query.Preload("Category").First(&product, "id = ?", product.ID).Error

	if err == nil {
		tx.Commit()
	}

	return product
}

func (repository *ProductRepositoryImpl) FindAll(ctx *fiber.Ctx, DB *gorm.DB) []domain.Product {

	products := []domain.Product{}

	keyword := ctx.Query("keyword")
	category_id := ctx.Query("category_id")

	tx := DB.Begin()
	defer tx.Rollback()

	query := tx.Preload("Category")

	if keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}
	if category_id != "" {
		query = query.Where("category_id = ?", category_id)
	}

	err := query.Find(&products).Error

	if err == nil {
		tx.Commit()
	}

	return products
}
