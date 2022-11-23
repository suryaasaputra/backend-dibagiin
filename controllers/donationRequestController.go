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

type donationRequestController struct {
	DonationRequestRepository repository.IDonationRequestRepository
}

func NewDonationRequestController(dr repository.IDonationRequestRepository) *donationRequestController {
	return &donationRequestController{
		DonationRequestRepository: dr,
	}
}

func (d donationRequestController) Create(ctx *gin.Context) {
	donationId := ctx.Param("donationId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := fmt.Sprintf("%v", userData["id"])

	request := models.CreateDonationRequestRequest{}

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

	donationRequest := models.DonationRequest{
		UserID:     userId,
		DonationID: donationId,
		Message:    request.Message,
	}

	resp, err := d.DonationRequestRepository.Create(donationRequest)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusBadRequest, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Send Request Success", resp)
	ctx.JSON(http.StatusCreated, response)
}

func (d donationRequestController) GetAll(ctx *gin.Context) {
	result, err := d.DonationRequestRepository.GetAll()
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "List of Requested Donation", result)
	ctx.JSON(http.StatusOK, response)
}

func (d donationRequestController) GetById(ctx *gin.Context) {
	id := ctx.Param("donationRequestId")
	result, err := d.DonationRequestRepository.GetById(id)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusNotFound, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Success", result)
	ctx.JSON(http.StatusOK, response)
}
