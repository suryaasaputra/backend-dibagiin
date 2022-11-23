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

func (d donationController) GetAll(ctx *gin.Context) {
	availableDonation := ctx.Query("available")

	if availableDonation == "true" {
		result, err := d.DonationRepository.GetAllAvailable()
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "List of Available Donations", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		result, err := d.DonationRepository.GetAll()
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "List of Donations", result)
		ctx.JSON(http.StatusOK, response)
	}

}

func (d donationController) GetDonationById(ctx *gin.Context) {
	donationId := ctx.Param("donationId")
	result, err := d.DonationRepository.GetById(donationId)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusNotFound, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Success", result)
	ctx.JSON(http.StatusOK, response)
}

func (d donationController) Edit(ctx *gin.Context) {
	donationId := ctx.Param("donationId")
	request := models.EditDonationRequest{}
	contentType := helpers.GetRequestHeaders(ctx).ContentType
	if contentType == "application/json" {
		err := ctx.ShouldBindJSON(&request)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		err := ctx.ShouldBind(&request)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "error binding request", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	}

	result, err := d.DonationRepository.Edit(donationId, request)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusBadRequest, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.GetResponse(false, http.StatusOK, "Success Update Donation Data", result)
	ctx.JSON(http.StatusOK, response)
}

func (d donationController) Delete(ctx *gin.Context) {
	donationId := ctx.Param("donationId")
	err := d.DonationRepository.Delete(donationId)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Success Delete Donation", nil)
	ctx.JSON(http.StatusOK, response)
}
