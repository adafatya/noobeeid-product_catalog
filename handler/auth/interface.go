package auth

import (
	"github.com/adafatya/noobeeid-product_catalog/constant"
	"github.com/adafatya/noobeeid-product_catalog/dto"
	"github.com/gofiber/fiber/v2"
)

type AuthUseCaseInterface interface {
	// create new auth
	// param email string, password string
	// return error
	CreateAuth(email, password string) *constant.Error

	// login
	// param email string, password string
	// return auth role, jwt token, error
	Login(email, password string) (dto.AuthLoginResponsePayloadDTO, *constant.Error)
}

type AuthHandlerInterface interface {
	// register
	Register(c *fiber.Ctx) error

	// login
	Login(c *fiber.Ctx) error
}
