package service

import (
	"test-golang-muhamad-isro/entity/domain"
	web "test-golang-muhamad-isro/entity/web/product"
	"test-golang-muhamad-isro/helper"
	repo_category_product "test-golang-muhamad-isro/repository/category_product"
	repo_product "test-golang-muhamad-isro/repository/product"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	ProductRepository         repo_product.ProductRepository
	CategoryProductRepository repo_category_product.CategoryProductRepository
	DB                        *gorm.DB
	Validate                  *validator.Validate
}

func NewProductService(productRepository repo_product.ProductRepository, DB *gorm.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Insert(ctx *fiber.Ctx, req web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	product := domain.Product{
		Name:        req.Name,
		CategoryID:  req.CategoryId,
		Description: req.Description,
	}

	product = service.ProductRepository.Insert(service.DB, product)

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAll(ctx *fiber.Ctx) []web.ProductResponse {

	products := service.ProductRepository.FindAll(ctx, service.DB)

	return helper.ToProductResponses(products)
}
