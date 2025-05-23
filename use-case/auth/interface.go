package auth

import (
	"github.com/adafatya/noobeeid-product_catalog/constant"
	"github.com/adafatya/noobeeid-product_catalog/dto"
	"github.com/adafatya/noobeeid-product_catalog/entity"
)

type AuthRepositoryInterface interface {
	// store new auth
	// param auth entity
	// return error
	Store(auth entity.Auth) error

	// get auth by email
	// param email string
	// return auth entity, error
	GetByEmail(email string) (entity.Auth, error)

	// set auth jwt token to redis
	// param id int, token string
	// return error
	SetAuthJWTToken(id int, token string) error

	// get auth jwt token from redis
	// param id int
	// return token string, error
	GetAuthJWTToken(id int) (string, error)
}

type AuthUseCaseInterface interface {
	// create new auth
	// param email string, password string
	// return error
	CreateAuth(email, password string) *constant.Error

	// login
	// param email string, password string
	// return login response payload dto, error
	Login(email, password string) (dto.AuthLoginResponsePayloadDTO, *constant.Error)
}
