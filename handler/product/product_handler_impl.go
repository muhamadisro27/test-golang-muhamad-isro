package handler

import (
	"fmt"
	"net/http"
	"test-golang-muhamad-isro/entity/domain"
	webResponse "test-golang-muhamad-isro/entity/web"
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
	// categoryCreateRequest := webRequest.ProductCreateRequest{}

	product := domain.Product{}

	fmt.Println(ctx.Body)
	// // helper.ReadFromRequestBody(ctx.Body, &categoryCreateRequest)

	// categoryResponse := handler.ProductService.Insert(ctx, categoryCreateRequest)

	// webResponse := webResponse.WebResponse{
	// 	Code:   http.StatusOK,
	// 	Status: "OK",
	// 	Data:   categoryResponse,
	// }

	// helper.WriteToJsonResponse(w, webResponse)

	return ctx.Status(201).JSON(product)
}

func (handler *ProductHandlerImpl) FindAll(ctx *fiber.Ctx) error {

	webResponse := webResponse.WebResponse{
		Code: http.StatusOK,
		Data: productResponses,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}
