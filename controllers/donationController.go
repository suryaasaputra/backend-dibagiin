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
		response := helpers.GetResponse(true, http.StatusBadRequest, "request tidak valid", nil)
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
		Weight:      request.Weight,
		Lat:         request.Lat,
		Lng:         request.Lng,
		PhotoUrl:    photoURL,
		Location:    request.Location,
	}

	resp, err := d.DonationRepository.Create(newDonation)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusBadRequest, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.GetResponse(false, http.StatusOK, "Berhasil membuat donasi", resp)
	ctx.JSON(http.StatusCreated, response)
}

func (d donationController) GetAll(ctx *gin.Context) {
	availability := ctx.Query("available")
	location := ctx.Query("location")
	title := ctx.Query("title")

	if availability == "true" {
		result, err := d.DonationRepository.GetAllAvailable()
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "Berhasil mendapatkan daftar donasi", result)
		ctx.JSON(http.StatusOK, response)
	} else if location != "" {
		result, err := d.DonationRepository.GetAllByLocation(location)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "Berhasil mendapatkan daftar donasi", result)
		ctx.JSON(http.StatusOK, response)
	} else if title != "" {
		result, err := d.DonationRepository.GetAllByKeyword(title)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "Berhasil mendapatkan daftar donasi", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		result, err := d.DonationRepository.GetAll()
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "Berhasil mendapatkan daftar donasi", result)
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
			response := helpers.GetResponse(true, http.StatusBadRequest, "request tidak valid", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
	} else {
		err := ctx.ShouldBind(&request)
		if err != nil {
			response := helpers.GetResponse(true, http.StatusBadRequest, "request tidak valid", nil)
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

	response := helpers.GetResponse(false, http.StatusOK, "Berhasil ubah data donasi", result)
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
	response := helpers.GetResponse(false, http.StatusOK, "Berhasil hapus donasi", nil)
	ctx.JSON(http.StatusOK, response)
}
