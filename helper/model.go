package helper

import (
	"test-golang-muhamad-isro/entity/domain"
	"test-golang-muhamad-isro/entity/web/product"
)

func ToProductResponse(product domain.Product) web.ProductResponse {
	return web.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
	}
}

func ToProductResponses(products []domain.Product) []web.ProductResponse {

	var productResponses []web.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}

	return productResponses
}
