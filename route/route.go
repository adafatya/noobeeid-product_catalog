package route

import "github.com/gofiber/fiber/v2"

type RouterInterface interface {
	Route(app *fiber.App)
}

func Route(
	app *fiber.App,
	authRouter RouterInterface,
	categoryRouter RouterInterface,
) {
	authRouter.Route(app)
	categoryRouter.Route(app)
}
