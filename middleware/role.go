package middleware

import (
	"net/http"

	"github.com/aldysp34/sm_padang/dto"
	"github.com/gin-gonic/gin"
)

func IsRole(allowedRoles ...uint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, role := range allowedRoles {
			if ctx.Value("role_id") == role {
				ctx.Next()
				return
			}
		}
		ctx.AbortWithStatusJSON(http.StatusForbidden, dto.Response{Message: "role doesnt have access"})
	}
}
