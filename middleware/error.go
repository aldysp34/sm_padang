package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/aldysp34/sm_padang/apperr"
	"github.com/aldysp34/sm_padang/dto"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				c.AbortWithStatusJSON(http.StatusGatewayTimeout, dto.Response{Message: "request timeout"})
				return
			}
			switch e := err.Err.(type) {
			case *apperr.CustomError:
				c.AbortWithStatusJSON(e.Code, e.ConvertToErrorResponse())
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			}
		}
	}
}
