package controllers

import (
	"dibagi/helpers"
	"dibagi/repository"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type donationHistoryController struct {
	DonationHistoryRepository repository.IDonationHistoryRepository
}

func NewDonationHistoryController(dr repository.IDonationHistoryRepository) *donationHistoryController {
	return &donationHistoryController{
		DonationHistoryRepository: dr,
	}
}

func (d donationHistoryController) GetAllByUserId(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := fmt.Sprintf("%v", userData["id"])
	result, err := d.DonationHistoryRepository.GetAllByUserId(userId)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "List of Claimed Donation", result)
	ctx.JSON(http.StatusOK, response)
}
