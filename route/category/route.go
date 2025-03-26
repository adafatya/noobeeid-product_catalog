package category

import "github.com/gofiber/fiber/v2"

type CategoryRouter struct {
	h CategoryHandlerInterface
}

func NewCategoryRouter(h CategoryHandlerInterface) CategoryRouter {
	return CategoryRouter{
		h: h,
	}
}

func (r CategoryRouter) Route(app *fiber.App) {
	v1 := app.Group("/v1/categories")

	v1.Get("/", r.h.GetAllCategory)
}
