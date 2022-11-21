package middlewares

import (
	"dibagi/helpers"
	"dibagi/repository"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type userMiddleware struct {
	UserRepository repository.IUserRepository
}

func NewUserMiddleware(userRepository repository.IUserRepository) *userMiddleware {
	return &userMiddleware{
		UserRepository: userRepository,
	}
}
func (u userMiddleware) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}

func (u userMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userNameURL := ctx.Param("userName")

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userName := fmt.Sprintf("%v", userData["user_name"])

		if userNameURL != userName {
			response := helpers.GetResponse(true, http.StatusUnauthorized, "You are not allowed to Access this data", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Next()
	}
}
