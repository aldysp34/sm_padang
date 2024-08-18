package handler

import (
	"net/http"
	"strconv"

	"github.com/aldysp34/sm_padang/apperr"
	"github.com/aldysp34/sm_padang/dto"
	"github.com/aldysp34/sm_padang/logger"
	"github.com/aldysp34/sm_padang/usecase"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUc *usecase.AdminUsecase
}

func NewAdminHandler(admin *usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUc: admin,
	}
}

func (ah *AdminHandler) ApproveRejectRequest(c *gin.Context) {
	var req dto.ReqApproval

	if err := c.BindQuery(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	if err := ah.adminUc.ApproveRejectRequest(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{Message: "successfully change request status"})

}

func (ah *AdminHandler) GetAllBarang(c *gin.Context) {
	res, err := ah.adminUc.GetAllBarang(c.Request.Context())
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully get all barang"})
}

func (ah *AdminHandler) GetAllBarangIn(c *gin.Context) {
	res, err := ah.adminUc.GetAllBarangIn(c.Request.Context())
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully get all barang in"})
}

func (ah *AdminHandler) DeleteBarangIn(c *gin.Context) {
	var req dto.ReqNewBarang
	if err := c.BindQuery(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.DeleteBarangIn(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully delete barang in"})
}

func (ah *AdminHandler) GetAllBarangOut(c *gin.Context) {
	res, err := ah.adminUc.GetAllBarangOut(c.Request.Context())
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully get all barang out"})
}

func (ah *AdminHandler) CreateNewBarang(c *gin.Context) {
	var req dto.ReqNewBarang
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.CreateNewBarang(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully create new barang"})
}

func (ah *AdminHandler) EditBarang(c *gin.Context) {
	var req dto.ReqNewBarang
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	var req2 dto.ReqNewSupplier
	if err := c.BindQuery(&req2); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	req.Id = req2.Id
	if err := ah.adminUc.EditBarang(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully edit barang"})
}

func (ah *AdminHandler) DeleteBarang(c *gin.Context) {
	var req dto.ReqNewBarang
	if err := c.BindQuery(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.DeleteBarang(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully delete barang"})
}

func (ah *AdminHandler) GetAllUser(c *gin.Context) {
	res, err := ah.adminUc.GetAllUser(c.Request.Context())
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully get all users"})
}

func (ah *AdminHandler) CreateNewUser(c *gin.Context) {
	var req dto.ReqNewUser
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.CreateNewUser(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully create new user"})
}

func (ah *AdminHandler) EditUser(c *gin.Context) {
	var req dto.ReqNewUser
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	var req2 dto.ReqNewUser
	if err := c.BindQuery(&req2); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	req.Id = req2.Id
	res, err := ah.adminUc.EditUser(c.Request.Context(), req)
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully edit user"})
}

func (ah *AdminHandler) DeleteUser(c *gin.Context) {
	var req dto.ReqNewUser
	if err := c.BindQuery(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.DeleteUser(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully delete user"})
}

func (ah *AdminHandler) CreateNewBrand(c *gin.Context) {
	var req dto.ReqNewBrandSatuan
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.CreateNewBrand(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully create new brand"})
}

func (ah *AdminHandler) GetAllBrand(c *gin.Context) {
	res, err := ah.adminUc.GetAllBrand(c.Request.Context())
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully get all brands"})
}

func (ah *AdminHandler) EditBrand(c *gin.Context) {
	var req dto.ReqNewBrandSatuan
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	var req2 dto.ReqNewSupplier
	if err := c.BindQuery(&req2); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	req.Id = req2.Id
	if err := ah.adminUc.EditBrand(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully edit brand"})
}

func (ah *AdminHandler) DeleteBrand(c *gin.Context) {
	var req dto.ReqNewBrandSatuan
	if err := c.BindQuery(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.DeleteBrand(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully delete brand"})
}

func (ah *AdminHandler) CreateNewSatuan(c *gin.Context) {
	var req dto.ReqNewBrandSatuan
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.CreateNewSatuan(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully create new satuan"})
}

func (ah *AdminHandler) GetAllSatuan(c *gin.Context) {
	res, err := ah.adminUc.GetAllSatuan(c.Request.Context())
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully get all satuan"})
}

func (ah *AdminHandler) EditSatuan(c *gin.Context) {
	var req dto.ReqNewBrandSatuan
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	var req2 dto.ReqNewBrandSatuan
	if err := c.BindQuery(&req2); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	req.Id = req2.Id
	if err := ah.adminUc.EditSatuan(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully edit satuan"})
}

func (ah *AdminHandler) DeleteSatuan(c *gin.Context) {
	var req dto.ReqNewBrandSatuan
	if err := c.BindQuery(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.DeleteSatuan(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully delete satuan"})
}

func (ah *AdminHandler) CreateNewSupplier(c *gin.Context) {
	var req dto.ReqNewSupplier
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.CreateNewSupplier(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully create new supplier"})
}

func (ah *AdminHandler) GetAllSupplier(c *gin.Context) {
	res, err := ah.adminUc.GetAllSupplier(c.Request.Context())
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully get all suppliers"})
}

func (ah *AdminHandler) EditSupplier(c *gin.Context) {
	var req dto.ReqNewSupplier
	if err := c.BindJSON(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	var req2 dto.ReqNewSupplier
	if err := c.BindQuery(&req2); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	req.Id = req2.Id
	if err := ah.adminUc.EditSupplier(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully edit supplier"})
}

func (ah *AdminHandler) DeleteSupplier(c *gin.Context) {
	var req dto.ReqNewSupplier
	if err := c.BindQuery(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.DeleteSupplier(c.Request.Context(), req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Message: "successfully delete supplier"})
}

func (ah *AdminHandler) GetAllRequest(c *gin.Context) {
	res, err := ah.adminUc.GetAllRequest(c.Request.Context())
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "successfully get all requests"})
}

func (ah *AdminHandler) CreateBarangIn(c *gin.Context) {
	var data dto.ReqNewBarangIn
	if err := c.BindJSON(&data); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	if err := ah.adminUc.CreateBarangIn(c, data); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, dto.Response{Message: "successfully create barang in"})
}

func (ah *AdminHandler) GetBarangByID(c *gin.Context) {
	var req dto.ReqNewBarang
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Log.Errorf(apperr.ErrInvalidID.Error())
		c.Error(apperr.ErrInvalidID)
		return
	}

	req.Id = uint(id)
	res, err := ah.adminUc.GetBarangByID(c, req)
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dto.Response{Data: res, Message: "success get barang by id"})
}

func (ah *AdminHandler) DownloadLaporan(c *gin.Context) {
	var req dto.ReqDate
	if err := c.BindQuery(&req); err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	file, err := ah.adminUc.DownloadXLSX(c, req.StartDate, req.EndDate)
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	buf, err := file.WriteToBuffer()
	if err != nil {
		logger.Log.Errorf(err.Error())
		c.Error(err)
		return
	}

	// now := time.Now().Format("2006-01-02")
	fileName := "attachment; filename=laporan.xlsx"

	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Header("Content-Disposition", fileName)
	c.Header("Content-Length", strconv.Itoa(len(buf.Bytes())))
	c.Writer.Write(buf.Bytes())

}
