package handler

import (
	"encoding/json"
	"net/http"
	"test-golang-muhamad-isro/entity/domain"
	webResponse "test-golang-muhamad-isro/entity/web"
	webRequest "test-golang-muhamad-isro/entity/web/product"
	"test-golang-muhamad-isro/helper"
	service "test-golang-muhamad-isro/service/product"

	"github.com/gofiber/fiber/v2"
)

type ProductHandlerImpl struct {
	ProductService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return &ProductHandlerImpl{
		ProductService: productService,
	}
}

func (handler *ProductHandlerImpl) Insert(ctx *fiber.Ctx) error {
	product := new(domain.Product)
	err := json.Unmarshal(ctx.Body(), product)

	helper.PanicIfError(err)

	productRequest := webRequest.ProductCreateRequest{
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.CategoryID,
	}

	productResponse := handler.ProductService.Insert(ctx, productRequest)

	webResponse := webResponse.WebResponse{
		Code: http.StatusOK,
		Data: productResponse,
	}

	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func (handler *ProductHandlerImpl) FindAll(ctx *fiber.Ctx) error {

	productResponses := handler.ProductService.FindAll(ctx)

	webResponse := webResponse.WebResponse{
		Code: http.StatusOK,
		Data: productResponses,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}
