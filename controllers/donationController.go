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

// Create donation godoc
// @Summary      Create Donation
// @Description  Create new donation
// @Param donation body models.CreateDonationRequest true "Donation data"
// @Tags         Donation
// @Accept       mpfd
// @Produce      json
// @Success      201  {object}  helpers.Response{data=models.CreateDonationResponse}
// @Router       /donation [post]
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

// Get All donation godoc
// @Summary      Get All Donation
// @Description  Get list donation
// @Param        available   path  bool  false  "availability"
// @Param        keyword   path  string  false  "keyword"
// @Tags         Donation
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response{data=[]models.GetDonationsResponse}
// @Router       /donation [get]
func (d donationController) GetAll(ctx *gin.Context) {
	availability := ctx.Query("available")
	keyword := ctx.Query("keyword")

	if availability == "true" {
		result, err := d.DonationRepository.GetAllAvailable()
		if err != nil {
			response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		response := helpers.GetResponse(false, http.StatusOK, "Berhasil mendapatkan daftar donasi", result)
		ctx.JSON(http.StatusOK, response)
	} else if keyword != "" {
		result, err := d.DonationRepository.GetAllByKeyword(keyword)
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

// Get donation detail godoc
// @Summary      Get Donation Detail
// @Description  Get donation detail by id
// @Param        donation_id   path  string  true  "Donation ID"
// @Tags         Donation
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response{data=models.GetDonationsResponse}
// @Router       /donation/{donation_id} [get]
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

// Edit donation detail godoc
// @Summary      Edit Donation Detail
// @Description  Edit donation detail by id
// @Param        donation body  models.EditDonationRequest  true  "Donation data"
// @Param        donation_id   path  string  true  "Donation ID"
// @Tags         Donation
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response{data=models.EditDonationResponse}
// @Router       /donation/{donation_id} [put]
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

// Delete donation  godoc
// @Summary      Delete Donation
// @Description  Delete donation  by id
// @Param        donation_id   path  string  true  "Donation ID"
// @Tags         Donation
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response
// @Router       /donation/{donation_id} [delete]
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
