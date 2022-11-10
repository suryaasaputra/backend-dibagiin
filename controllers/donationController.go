package controllers

import (
	"dibagi/helpers"
	"dibagi/models"
	"dibagi/repository"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type donationController struct {
	DonationRepository repository.IDonationRepository
}

func NewDonationController(dr repository.IDonationRepository) *donationController {
	return &donationController{
		DonationRepository: dr,
	}
}

func (d donationController) Create(ctx *gin.Context) {
	request := models.CreateDonationRequest{}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	id := fmt.Sprintf("%v", userData["id"])
	err := ctx.ShouldBind(&request)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	photoURL, err := helpers.UploadToBucket(ctx.Request, "donation_photo", id)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	newDonation := models.Donation{
		UserID:      id,
		Title:       request.Title,
		Description: request.Description,
		PhotoUrl:    photoURL,
		Location:    request.Location,
	}

	resp, err := d.DonationRepository.Create(newDonation)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusBadRequest, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.GetResponse(false, http.StatusOK, "Create Donation Success", resp)
	ctx.JSON(http.StatusCreated, response)
}
