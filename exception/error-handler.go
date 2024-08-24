package exception

import (
	"test-golang-muhamad-isro/entity/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// If notFoundError
	if notFoundError(ctx, err) {
		return nil
	}

	// If validationError
	if validationError(ctx, err) {
		return nil
	}

	// If Internal Server Error
	return InternalServerError(ctx, err)
}

func validationError(ctx *fiber.Ctx, err error) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		erroResponse := web.ErrorResponse{
			Errors: exception.Error(),
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(erroResponse) == nil
	}
	return false
}

func notFoundError(ctx *fiber.Ctx, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		errorResponse := web.ErrorResponse{
			Errors: exception.Error,
		}

		return ctx.Status(fiber.StatusNotFound).JSON(errorResponse) == nil
	}
	return false
}

func InternalServerError(ctx *fiber.Ctx, err error) error {
	errorResponse := web.ErrorResponse{
		Errors: err.Error(),
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(errorResponse)
}
