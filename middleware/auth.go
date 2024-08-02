package middleware

import (
	"context"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/aldysp34/sm_padang/apperr"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if gin.Mode() == gin.TestMode {
			return
		}

		// session := sessions.Default(ctx)

		for pathPattern, allowedMethods := range common.AllowedRoles {
			matched, _ := regexp.MatchString(pathPattern, ctx.Request.URL.Path)
			if matched {
				if util.Contains(allowedMethods, ctx.Request.Method) {
					ctx.Next()
					return
				}
			}
		}
		var response dto.Response

		header := ctx.Request.Header["Authorization"]
		if len(header) == 0 {
			response.Message = apperr.ErrBearerTokenInvalid.Error()
			ctx.AbortWithStatusJSON(apperr.ErrBearerTokenInvalid.Code, response)
			return
		}

		splittedHeader := strings.Split(header[0], " ")
		if len(splittedHeader) != 2 {
			response.Message = apperr.ErrUnauthorize.Error()
			ctx.AbortWithStatusJSON(apperr.ErrUnauthorize.Code, response)
			return
		}

		claims := &dto.JWTClaims{}

		token, err := jwt.ParseWithClaims(splittedHeader[1], claims, func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, apperr.ErrWrongCredentials
			}
			return []byte(os.Getenv("API_SECRET")), nil
		})
		if err != nil {
			response.Message = err.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		_, ok := token.Claims.(*dto.JWTClaims)
		if !ok || !token.Valid {
			response.Message = apperr.ErrUnauthorize.Error()
			ctx.AbortWithStatusJSON(apperr.ErrUnauthorize.Code, response)
			return
		}

		newCtx := context.WithValue(ctx.Request.Context(), common.ID, claims.ID)
		ctx.Request = ctx.Request.WithContext(newCtx)
		ctx.Next()
	}
}
