package dto

import (
	"time"
)

type ReqNewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReqUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type VerifyResponse struct {
	Code string `json:"code"`
}

type ChangePassword struct {
	Email       string `json:"email"`
	NewPassword string `json:"password,omitempty"`
}

type ForgotPassword struct {
	Token      string `json:"token,omitempty"`
	Email      string `json:"email"`
	Expiration time.Time
}
