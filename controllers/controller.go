package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserController            *userController
	DonationController        *donationController
	DonationRequestController *donationRequestController
}

func NewController(userController *userController, donationController *donationController, dondonationRequestController *donationRequestController) Controller {
	return Controller{
		UserController:            userController,
		DonationController:        donationController,
		DonationRequestController: dondonationRequestController,
	}
}

func (c Controller) HomeController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}
