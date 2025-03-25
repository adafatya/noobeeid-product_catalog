package auth

import (
	"net/http"
	"regexp"

	"github.com/adafatya/noobeeid-product_catalog/constant"
	"github.com/adafatya/noobeeid-product_catalog/dto"
	"github.com/adafatya/noobeeid-product_catalog/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	uc AuthUseCaseInterface
}

func NewAuthHandler(uc AuthUseCaseInterface) AuthHandler {
	return AuthHandler{
		uc: uc,
	}
}

func validateEmail(email string) *constant.Error {
	// empty string validation
	if email == "" {
		return constant.ErrRequiredEmail
	}
	// email regex validation
	match, err := regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	if err != nil {
		return constant.ErrUnknownError
	}
	if !match {
		return constant.ErrInvalidEmail
	}

	return nil
}

func validatePassword(password string) *constant.Error {
	// empty string validation
	if password == "" {
		return constant.ErrRequiredPassword
	}
	// password length validation
	if len(password) < 6 {
		return constant.ErrInvalidPassword
	}

	return nil
}

func (h AuthHandler) Register(c *fiber.Ctx) error {
	auth := new(dto.AuthRegisterRequestDTO)

	// parse body to auth register request dto
	err := c.BodyParser(auth)
	if err != nil {
		return utils.ErrorResponse(c, *constant.ErrUnknownError)
	}

	// email validation
	valErr := validateEmail(auth.Email)
	if valErr != nil && valErr == constant.ErrUnknownError {
		return utils.ErrorResponse(c, *valErr)
	}
	if valErr != nil {
		return utils.ErrorResponse(c, *valErr)
	}

	// password validation
	valErr = validatePassword(auth.Password)
	if valErr != nil {
		return utils.ErrorResponse(c, *valErr)
	}

	// create auth
	ucErr := h.uc.CreateAuth(auth.Email, auth.Password)
	if ucErr != nil && ucErr == constant.ErrDuplicateEmail {
		return utils.ErrorResponse(c, *ucErr)
	}
	if ucErr != nil {
		return utils.ErrorResponse(c, *ucErr)
	}

	return utils.SuccessResponse(c, http.StatusCreated, "registration success")
}

func (h AuthHandler) Login(c *fiber.Ctx) error {
	auth := new(dto.AuthLoginRequestDTO)

	// parse body to auth login request dto
	err := c.BodyParser(auth)
	if err != nil {
		return utils.ErrorResponse(c, *constant.ErrUnknownError)
	}

	// validate email
	valErr := validateEmail(auth.Email)
	if valErr != nil && valErr == constant.ErrUnknownError {
		return utils.ErrorResponse(c, *valErr)
	}
	if valErr != nil {
		return utils.ErrorResponse(c, *valErr)
	}

	// validate password
	valErr = validatePassword(auth.Password)
	if valErr != nil {
		return utils.ErrorResponse(c, *valErr)
	}

	// login
	payload, ucErr := h.uc.Login(auth.Email, auth.Password)
	if ucErr != nil && (ucErr == constant.ErrInvalidEmail || ucErr == constant.ErrInvalidPassword) {
		return utils.ErrorResponse(c, *ucErr)
	}
	if ucErr != nil {
		return utils.ErrorResponse(c, *ucErr)
	}

	return utils.SuccessResponseWithPayload(c, http.StatusOK, "login success", payload)
}
