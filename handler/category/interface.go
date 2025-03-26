package category

import (
	"github.com/adafatya/noobeeid-product_catalog/constant"
	"github.com/adafatya/noobeeid-product_catalog/entity"
	"github.com/gofiber/fiber/v2"
)

type CategoryUseCaseInterface interface {
	// get all category
	//
	// return list of category, error
	GetAllCategory() ([]entity.Category, *constant.Error)
}

type CategoryHandlerInterface interface {
	// get all category
	GetAllCategory(c *fiber.Ctx) error
}
