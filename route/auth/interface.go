package auth

import "github.com/gofiber/fiber/v2"

type AuthHandlerInterface interface {
	// register
	Register(c *fiber.Ctx) error

	// login
	Login(c *fiber.Ctx) error
}
