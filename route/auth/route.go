package auth

import "github.com/gofiber/fiber/v2"

type AuthRouter struct {
	h AuthHandlerInterface
}

func NewAuthRouter(h AuthHandlerInterface) AuthRouter {
	return AuthRouter{
		h: h,
	}
}

func (r AuthRouter) Route(app *fiber.App) {
	v1 := app.Group("/v1/auth")
	v1.Post("/register", r.h.Register)
	v1.Post("/login", r.h.Login)
}
