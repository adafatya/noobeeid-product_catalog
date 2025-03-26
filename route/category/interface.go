package category

import "github.com/gofiber/fiber/v2"

type CategoryHandlerInterface interface {
	// get all category
	GetAllCategory(c *fiber.Ctx) error
}
