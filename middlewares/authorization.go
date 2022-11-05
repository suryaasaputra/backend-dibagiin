package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userNameURL := ctx.Param("userName")

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userName := fmt.Sprintf("%v", userData["user_name"])

		if userNameURL != userName {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": "You are not allowed to Edit this data",
			})
			return
		}
	}
}
