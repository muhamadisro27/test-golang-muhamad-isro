package repository

import (
	"test-golang-muhamad-isro/entity/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryProductRepository interface {
	FindById(DB *gorm.DB, category_id uuid.UUID) (domain.ProductCategory, error)
}
