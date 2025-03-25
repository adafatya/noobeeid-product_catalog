package utils

import (
	"github.com/adafatya/noobeeid-product_catalog/constant"
	"github.com/gofiber/fiber/v2"
)

// fiber success response
// param c fiber context, http status code int, message string
// return error
func SuccessResponse(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": true,
		"message": msg,
	})
}

// fiber success response with payload
// param c fiber context, http status code int, message string, payload any
// return error
func SuccessResponseWithPayload(c *fiber.Ctx, status int, msg string, payload any) error {
	return c.Status(status).JSON(fiber.Map{
		"success": true,
		"message": msg,
		"payload": payload,
	})
}

// fiber error response
// param c fiber context, error
// return error
func ErrorResponse(c *fiber.Ctx, err constant.Error) error {
	return c.Status(err.HttpStatus).JSON(fiber.Map{
		"success":    false,
		"message":    err.Message,
		"error":      err.ErrorMessage,
		"error_code": err.ErrorCode,
	})
}
