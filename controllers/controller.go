package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserController     *userController
	DonationController *donationController
}

func NewController(userController *userController, donationController *donationController) Controller {
	return Controller{
		UserController:     userController,
		DonationController: donationController,
	}
}

func (c Controller) HomeController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}
