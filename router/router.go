package router

import (
	handler "test-golang-muhamad-isro/handler/product"

	"github.com/gofiber/fiber/v2"
)

func ProductRouter(app *fiber.App, productHandler handler.ProductHandler) {
	products := app.Group("/api/v1/products/")
	products.Get("/", productHandler.FindAll)
	products.Post("/", productHandler.Insert)
}
