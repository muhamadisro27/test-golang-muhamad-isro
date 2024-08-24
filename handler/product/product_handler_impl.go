package handler

import (
	"encoding/json"
	"net/http"
	"test-golang-muhamad-isro/entity/domain"
	webResponse "test-golang-muhamad-isro/entity/web"
	webRequest "test-golang-muhamad-isro/entity/web/product"
	service "test-golang-muhamad-isro/service/product"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

	var payload map[string]interface{}
	err := json.Unmarshal(ctx.Body(), &payload)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid request payload",
		})
	}

	categoryIDStr, ok := payload["category_id"].(string)
	if !ok || categoryIDStr == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": "category id is required and cannot be empty",
		})
	}

	categoryID, err := uuid.Parse(categoryIDStr)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": "Invalid Category ID format",
		})
	}

	product := new(domain.Product)
	err = json.Unmarshal(ctx.Body(), product)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"errors": "Failed to parse request body",
		})
	}

	product.CategoryID = categoryID

	productRequest := webRequest.ProductCreateRequest{
		Name:        product.Name,
		Description: product.Description,
		CategoryId:  product.CategoryID,
	}

	productResponse := handler.ProductService.Insert(ctx, productRequest)

	webResponse := webResponse.WebResponse{
		Code: http.StatusCreated,
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
