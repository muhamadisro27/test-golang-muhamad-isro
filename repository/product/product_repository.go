package repository

import (
	"test-golang-muhamad-isro/entity/domain"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Insert(DB *gorm.DB, product domain.Product) domain.Product
	FindAll(ctx *fiber.Ctx, DB *gorm.DB) []domain.Product
}
