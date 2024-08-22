package service

import (
	"test-golang-muhamad-isro/entity/web/product"

	"github.com/gofiber/fiber/v2"
)

type ProductService interface {
	Insert(ctx *fiber.Ctx, req web.ProductCreateRequest) web.ProductResponse
	FindAll(ctx *fiber.Ctx) []web.ProductResponse
}
