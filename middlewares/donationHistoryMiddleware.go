package middlewares

import (
	"dibagi/helpers"
	"dibagi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type donationHistoryMiddleware struct {
	DonationHistoryRepository repository.IDonationHistoryRepository
}

func NewDonationHistoryMiddleware(donationHistoryRepository repository.IDonationHistoryRepository) *donationHistoryMiddleware {
	return &donationHistoryMiddleware{
		DonationHistoryRepository: donationHistoryRepository,
	}
}

func (d donationHistoryMiddleware) CheckIfExist() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		donationRequestId := ctx.Param("donationRequestId")
		result, err := d.DonationHistoryRepository.GetAllByDonationRequestId(donationRequestId)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				ctx.Next()
				return
			}
			response := helpers.GetResponse(true, http.StatusInternalServerError, "Terjadi Kesalahan", nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}

		for _, v := range result {

			if v.DonationRequest.UserID == v.UserID {
				response := helpers.GetResponse(true, http.StatusBadRequest, "Anda sudah mengkonfirmasi / menolak permintaan ini", nil)
				ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
				return
			}
		}
		ctx.Next()
	}
}
