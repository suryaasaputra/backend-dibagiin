package controllers

import (
	"dibagi/helpers"
	"dibagi/repository"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type notificationController struct {
	NotificationRepository repository.INotificationRepository
}

func NewNotificationController(dr repository.INotificationRepository) *notificationController {
	return &notificationController{
		NotificationRepository: dr,
	}
}

func (d notificationController) GetAllByUserId(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := fmt.Sprintf("%v", userData["id"])
	result, err := d.NotificationRepository.GetAllByUserId(userId)
	if err != nil {
		response := helpers.GetResponse(true, http.StatusInternalServerError, err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := helpers.GetResponse(false, http.StatusOK, "Daftar donasi yang diminta dan diterima", result)
	ctx.JSON(http.StatusOK, response)
}
