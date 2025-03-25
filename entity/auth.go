package entity

import (
	"log"
	"strconv"

	"github.com/adafatya/noobeeid-product_catalog/utils"
)

type Auth struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewAuth(email, password string) Auth {
	return Auth{
		Email:    email,
		Password: password,
		Role:     "merchant",
	}
}

func (e *Auth) EncryptPassword() error {
	// create encrypted password
	pwd, err := utils.HashPassword(e.Password)
	if err != nil {
		log.Println("error saat enkripsi password: ", err)
		return err
	}

	// update auth password to encrypted password
	e.Password = pwd

	return nil
}

func (e Auth) VerifyPassword(password string) bool {
	// verify password
	return utils.VerifyPassword(password, e.Password)
}

func (e Auth) GenerateJWTToken() (string, error) {
	token, err := utils.GenerateJWTToken(strconv.Itoa(e.Id))
	if err != nil {
		log.Println("error saat pembuatan token jwt: ", err)
		return "", err
	}

	return token, nil
}
