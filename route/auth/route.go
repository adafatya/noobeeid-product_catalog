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
	auth := app.Group("/auth")
	auth.Post("/register", r.h.Register)
	auth.Post("/login", r.h.Login)
}
