package middlewares

import (
	"dibagi/helpers"
	"dibagi/repository"
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
			response := helpers.GetResponse(true, http.StatusUnauthorized, "You are not allowed to Access this data", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Next()
	}
}
func DonationAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		donationId := ctx.Param("donationId")

		db := repository.GetDB()
		donationRepository := repository.NewDonationRepository(db)
		result, err := donationRepository.GetDonationById(donationId)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, "Something went wrong", nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := fmt.Sprintf("%v", userData["id"])

		if userId != result.UserID {
			response := helpers.GetResponse(true, http.StatusUnauthorized, "You are not allowed to Access this data", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Next()
	}
}
