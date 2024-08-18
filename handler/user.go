package handler

import (
	"net/http"

	"github.com/aldysp34/sm_padang/apperr"
	"github.com/aldysp34/sm_padang/dto"
	"github.com/aldysp34/sm_padang/logger"
	"github.com/aldysp34/sm_padang/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUc *usecase.UserUsecase
}

func NewUserHandler(admin *usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUc: admin,
	}
}

func (uh *UserHandler) CreateNewRequest(c *gin.Context) {
	var req dto.ReqNewRequest
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	if err := uh.userUc.CreateNewRequest(c, req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dto.Response{Message: "successfully create new request"})
}

func (uh *UserHandler) GetUserRequest(c *gin.Context) {
	data, err := uh.userUc.GetUserRequest(c)
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "successfully get user request", Data: data})
}

func (uh *UserHandler) GetAllBarang(c *gin.Context) {
	data := uh.userUc.GetAllBarang(c)
	if len(data) < 1 {
		logger.Log.Errorf(apperr.NewCustomError(http.StatusBadRequest, "no barang data").Error())
		c.Error(apperr.NewCustomError(http.StatusBadRequest, "no barang data"))
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "successfully get all barang", Data: data})

}
