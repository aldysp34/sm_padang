package server

import (
	"net/http"

	"github.com/aldysp34/sm_padang/common"
	"github.com/aldysp34/sm_padang/handler"
	"github.com/aldysp34/sm_padang/logger"
	"github.com/aldysp34/sm_padang/middleware"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	AdminHandler *handler.AdminHandler
	AuthHandler  *handler.AuthHandler
	UserHandler  *handler.UserHandler
}

func NewRouter(opts RouterOpts) *gin.Engine {
	r := gin.New()
	r.ContextWithFallback = true

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.CorsHandler())
	r.Use(middleware.Logger(logger.NewLogger()))
	r.Use(middleware.AuthorizeHandler())
	r.Use(middleware.ErrorHandler())

	api := r.Group("/api")

	api.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "HELLO WORLD")
	})

	admin := api.Group("/admin")
	admin.GET("/barangin", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetAllBarangIn)
	admin.DELETE("/barangin", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.DeleteBarangIn)
	admin.POST("/barangin", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.CreateBarangIn)

	admin.GET("/barangout", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetAllBarangOut)

	admin.GET("/barang", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetAllBarang)
	admin.POST("/barang", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.CreateNewBarang)
	admin.PUT("/barang", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.EditBarang)
	admin.DELETE("/barang", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.DeleteBarang)
	admin.GET("/barang/:id", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetBarangByID)

	admin.GET("/user", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetAllUser)
	admin.POST("/user", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.CreateNewUser)
	admin.PUT("/user", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.EditUser)
	admin.DELETE("/user", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.DeleteUser)

	admin.GET("/brand", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetAllBrand)
	admin.POST("/brand", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.CreateNewBrand)
	admin.PUT("/brand", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.EditBrand)
	admin.DELETE("/brand", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.DeleteBrand)

	admin.GET("/satuan", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetAllSatuan)
	admin.POST("/satuan", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.CreateNewSatuan)
	admin.PUT("/satuan", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.EditSatuan)
	admin.DELETE("/satuan", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.DeleteSatuan)

	admin.GET("/supplier", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetAllSupplier)
	admin.POST("/supplier", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.CreateNewSupplier)
	admin.PUT("/supplier", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.EditSupplier)
	admin.DELETE("/supplier", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.DeleteSupplier)

	admin.POST("/request/approve-reject", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.ApproveRejectRequest)
	admin.GET("/request", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.GetAllRequest)
	admin.GET("/laporan", middleware.IsRole(common.RoleAdmin), opts.AdminHandler.DownloadLaporan)

	auth := api.Group("/auth")
	auth.POST("/login", opts.AuthHandler.Login)

	user := api.Group("/user")
	user.POST("/request", middleware.IsRole(common.RoleUser), opts.UserHandler.CreateNewRequest)
	user.GET("/request", middleware.IsRole(common.RoleUser), opts.UserHandler.GetUserRequest)
	user.GET("/barang", middleware.IsRole(common.RoleUser), opts.UserHandler.GetAllBarang)
	api.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "HELLO WORLD")
	})

	return r
}
