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

// Send Donation Request godoc
// @Summary      Send Donation Request
// @Description  Send request to claim donation
// @Param        donation_id   path  string  true  "Donation ID"
// @Param        message  body  models.CreateDonationRequestRequest  true  "Donation request body"
// @Tags         Donation Request
// @Accept       json
// @Produce      json
// @Success      201  {object}  helpers.Response{data=models.CreateDonationRequestResponse}
// @Router       /donation/{donation_id}/request [post]
func (d donationRequestController) Create(ctx *gin.Context) {
	donationId := ctx.Param("donationId")
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := fmt.Sprintf("%v", userData["id"])

	request := models.CreateDonationRequestRequest{}

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
	response := helpers.GetResponse(false, http.StatusOK, "Berhasil mengirim permintaan", resp)
	ctx.JSON(http.StatusCreated, response)
}

// Get All Submitted Donation Request godoc
// @Summary      Get All Submitted Donation Request
// @Description  Get all submitted donation request
// @Tags         Donation Request
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response{data=models.GetDonationRequestResponse}
// @Router       /donation/request [get]
func (d donationRequestController) GetAllByUserId(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := fmt.Sprintf("%v", userData["id"])
	result, err := d.DonationRequestRepository.GetAllByUserId(userId)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Daftar permintaan yang dikirim", result)
	ctx.JSON(http.StatusOK, response)
}

// Get All Recived Donation Request godoc
// @Summary      Get All Recived Donation Request
// @Description  Get all recived donation request
// @Tags         Donation Request
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response{data=models.GetDonationRequestResponse}
// @Router       /request [get]
func (d donationRequestController) GetAllByDonatorId(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := fmt.Sprintf("%v", userData["id"])

	result, err := d.DonationRequestRepository.GetAllByDonatorId(userId)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Daftar permintaan yang diterima", result)
	ctx.JSON(http.StatusOK, response)
}

func (d donationRequestController) GetAllByDonationId(ctx *gin.Context) {
	donationId := ctx.Param("donationId")
	result, err := d.DonationRequestRepository.GetAllByDonationId(donationId)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Daftar permintaan pada donasi ini", result)
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

// Confirm Donation Request godoc
// @Summary      Confirm Donation Request
// @Description  Confirm request to claim donation
// @Param        donationRequest_id   path  string  true  "Donation Request ID"
// @Tags         Donation Request
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response
// @Router       /request/{donationRequest_id} [post]
func (d donationRequestController) Confirm(ctx *gin.Context) {
	id := ctx.Param("donationRequestId")

	err := d.DonationRequestRepository.Confirm(id)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Konfirmasi permintaan berhasil", nil)
	ctx.JSON(http.StatusOK, response)
}

// Reject Donation Request godoc
// @Summary      Reject Donation Request
// @Description  Reject donation request
// @Param        donationRequest_id   path  string  true  "Donation Request ID"
// @Tags         Donation Request
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response
// @Router       /request/{donationRequest_id} [delete]
func (d donationRequestController) Reject(ctx *gin.Context) {
	id := ctx.Param("donationRequestId")

	err := d.DonationRequestRepository.Reject(id)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Tolak permintaan berhasil", nil)
	ctx.JSON(http.StatusOK, response)
}

// Cancel Send Donation Request godoc
// @Summary      Cancel Send Donation Request
// @Description  Cancel Send donation request
// @Param        donationRequest_id   path  string  true  "Donation Request ID"
// @Tags         Donation Request
// @Accept       json
// @Produce      json
// @Success      200  {object}  helpers.Response
// @Router       /donation/request/{donationRequest_id} [delete]
func (d donationRequestController) Delete(ctx *gin.Context) {
	id := ctx.Param("requestId")

	err := d.DonationRequestRepository.Delete(id)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Sukses batalkan permintaan", nil)
	ctx.JSON(http.StatusOK, response)
}
