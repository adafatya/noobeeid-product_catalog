package category

import (
	"net/http"

	"github.com/adafatya/noobeeid-product_catalog/utils"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	uc CategoryUseCaseInterface
}

func NewCategoryHandler(uc CategoryUseCaseInterface) CategoryHandler {
	return CategoryHandler{
		uc: uc,
	}
}

func (h CategoryHandler) GetAllCategory(c *fiber.Ctx) error {
	categories, err := h.uc.GetAllCategory()
	if err != nil {
		return utils.ErrorResponse(c, *err)
	}

	return utils.SuccessResponseWithPayload(c, http.StatusOK, "get categories success", categories)
}
