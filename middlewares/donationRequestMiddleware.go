package middlewares

import (
	"dibagi/helpers"
	"dibagi/repository"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type donationRequestMiddleware struct {
	DonationRequestRepository repository.IDonationRequestRepository
}

func NewDonationRequestMiddleware(donationRequestRepository repository.IDonationRequestRepository) *donationRequestMiddleware {
	return &donationRequestMiddleware{
		DonationRequestRepository: donationRequestRepository,
	}
}

func (d donationRequestMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		donationRequestId := ctx.Param("donationRequestId")
		result, err := d.DonationRequestRepository.GetById(donationRequestId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response := helpers.GetResponse(true, http.StatusNotFound, "Request not found", nil)
				ctx.AbortWithStatusJSON(http.StatusNotFound, response)
				return
			}
			response := helpers.GetResponse(true, http.StatusInternalServerError, "Something went wrong", nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := fmt.Sprintf("%v", userData["id"])

		if userId != result.DonatorID {
			response := helpers.GetResponse(true, http.StatusUnauthorized, "You are not allowed to access this data", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Next()

	}
}
func (d donationRequestMiddleware) CheckIfExist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		donationId := ctx.Param("donationId")
		fmt.Println("masuk")
		result, err := d.DonationRequestRepository.GetAllByDonationId(donationId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Println("masuk2")
				ctx.Next()
				return
			}
			fmt.Println("masuk3")
			response := helpers.GetResponse(true, http.StatusInternalServerError, "Something went wrong", nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		fmt.Println("masuk4")
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := fmt.Sprintf("%v", userData["id"])

		for _, v := range result {
			if userId == v.UserID {
				fmt.Println("masuk5")
				fmt.Println(userId)
				fmt.Println(v.UserID)
				response := helpers.GetResponse(true, http.StatusBadRequest, "Request already exist", nil)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
				return
			}
		}
		fmt.Println("masuk6")
		ctx.Next()
	}
}
