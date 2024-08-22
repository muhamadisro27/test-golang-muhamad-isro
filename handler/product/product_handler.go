package handler

import "github.com/gofiber/fiber/v2"

type ProductHandler interface {
	Insert(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}