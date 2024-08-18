package handler

import (
	"net/http"

	"github.com/aldysp34/sm_padang/dto"
	"github.com/aldysp34/sm_padang/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth *usecase.AuthUsecase
}

func NewAuthHandler(auth *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		auth: auth,
	}
}

func (a *AuthHandler) Login(c *gin.Context) {
	var req dto.ReqUser
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(err)
		return
	}

	token, err := a.auth.Login(c, req)
	if err != nil {
		c.Error(err)

		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "login successfully", Data: token})
}
