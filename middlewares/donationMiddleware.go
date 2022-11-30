package middlewares

import (
	"dibagi/helpers"
	"dibagi/repository"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type donationMiddleware struct {
	DonationRepository repository.IDonationRepository
}

func NewDonationMiddleware(donationRepository repository.IDonationRepository) *donationMiddleware {
	return &donationMiddleware{
		DonationRepository: donationRepository,
	}
}

func (d donationMiddleware) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		donationId := ctx.Param("donationId")
		result, err := d.DonationRepository.GetById(donationId)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, "Terjadi kesalahan", nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := fmt.Sprintf("%v", userData["id"])

		if userId != result.UserID {
			response := helpers.GetResponse(true, http.StatusUnauthorized, "Akses ditolak, Anda tidak diizinkan mengubah data", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Next()

	}
}

func (d donationMiddleware) CheckDonator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		donationId := ctx.Param("donationId")
		result, err := d.DonationRepository.GetById(donationId)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, "Terjadi kesalahan", nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		if result.Status != "Tersedia" {
			response := helpers.GetResponse(true, http.StatusBadRequest, "Donasi sudah diambil oleh pengguna lain", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userId := fmt.Sprintf("%v", userData["id"])

		if userId == result.UserID {
			response := helpers.GetResponse(true, http.StatusBadRequest, "Anda tidak bisa mengirim permintaan ke Donasi Anda sendiri", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		ctx.Next()
	}
}
