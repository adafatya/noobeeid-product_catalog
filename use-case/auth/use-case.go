package auth

import (
	"github.com/adafatya/noobeeid-product_catalog/constant"
	"github.com/adafatya/noobeeid-product_catalog/dto"
	"github.com/adafatya/noobeeid-product_catalog/entity"
	"github.com/redis/go-redis/v9"
)

type AuthUseCase struct {
	repo AuthRepositoryInterface
}

func NewAuthUseCase(repo AuthRepositoryInterface) AuthUseCase {
	return AuthUseCase{
		repo: repo,
	}
}

func (uc AuthUseCase) CreateAuth(email, password string) *constant.Error {
	// check duplicate email
	exist, err := uc.repo.GetByEmail(email)
	if err != nil {
		return constant.ErrRepositoryError
	}
	if (exist != entity.Auth{}) {
		return constant.ErrDuplicateEmail
	}

	// create new auth entity
	auth := entity.NewAuth(email, password)

	// encrypt password
	err = auth.EncryptPassword()
	if err != nil {
		return constant.ErrUnknownError
	}

	// insert to db
	err = uc.repo.Store(auth)
	if err != nil {
		return constant.ErrRepositoryError
	}

	return nil
}

func (uc AuthUseCase) Login(email, password string) (dto.AuthLoginResponsePayloadDTO, *constant.Error) {
	// get auth data by email
	auth, err := uc.repo.GetByEmail(email)
	if err != nil {
		return dto.AuthLoginResponsePayloadDTO{}, constant.ErrRepositoryError
	}

	// check unregistered auth
	if (auth == entity.Auth{}) {
		return dto.AuthLoginResponsePayloadDTO{}, constant.ErrInvalidEmail
	}

	// verify password
	if !auth.VerifyPassword(password) {
		return dto.AuthLoginResponsePayloadDTO{}, constant.ErrInvalidPassword
	}

	// check if token is on redis
	token, err := uc.repo.GetAuthJWTToken(auth.Id)
	if err != nil && err != redis.Nil {
		return dto.AuthLoginResponsePayloadDTO{}, constant.ErrRepositoryError
	}

	if err == redis.Nil {
		// generate jwt token
		token, err = auth.GenerateJWTToken()
		if err != nil {
			return dto.AuthLoginResponsePayloadDTO{}, constant.ErrUnknownError
		}

		// set token to redis
		err := uc.repo.SetAuthJWTToken(auth.Id, token)
		if err != nil {
			return dto.AuthLoginResponsePayloadDTO{}, constant.ErrRepositoryError
		}
	}

	resp := dto.AuthLoginResponsePayloadDTO{
		AccessToken: token,
		Role:        auth.Role,
	}
	return resp, nil
}
