package server

import (
	"net/http"

	"github.com/aldysp34/sm_padang/logger"
	"github.com/aldysp34/sm_padang/middleware"
	"github.com/gin-gonic/gin"
)

type RouterOpts struct {
	// AuthHandler    *handler.AuthHandler
	// ProductHandler *handler.ProductHandler
	// ArticleHandler *handler.ArticleHandler
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

	auth := api.Group("/auth")
	auth.POST("/login", opts.AuthHandler.Login)
	auth.POST("/register", opts.AuthHandler.Register)
	auth.POST("/logout", opts.AuthHandler.Logout)

	return r
}
