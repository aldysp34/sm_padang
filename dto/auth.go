package dto

import (
	"time"
)

type ReqNewUser struct {
	Id       uint   `json:"id" form:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type ReqUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	RoleID   uint   `json:"role_id"`
	Role     string `json:"role_name"`
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
