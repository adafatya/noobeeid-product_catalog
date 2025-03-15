package entity

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
