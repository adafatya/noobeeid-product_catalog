package dto

type AuthRegisterRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponsePayloadDTO struct {
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
}
