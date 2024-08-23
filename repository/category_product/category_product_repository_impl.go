package repository

import (
	"test-golang-muhamad-isro/entity/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryProductRepositoryImpl struct {
}

func NewProductRepository() CategoryProductRepository {
	return &CategoryProductRepositoryImpl{}
}

func (repository *CategoryProductRepositoryImpl) FindById(DB *gorm.DB, category_id uuid.UUID) (domain.ProductCategory, error) {

	var category domain.ProductCategory

	err := DB.First(&category, "id = ?", category_id).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
