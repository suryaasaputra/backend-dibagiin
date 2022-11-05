package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserController *userController
}

func NewController(userController *userController) Controller {
	return Controller{
		UserController: userController,
	}
}

func (c Controller) HomeController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}
