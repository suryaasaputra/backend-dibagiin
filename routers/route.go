package routers

import (
	"dibagi/controllers"
	"dibagi/middlewares"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(ctl controllers.Controller, mdl middlewares.Middleware) error {

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", ctl.HomeController)

	//manual login
	r.POST("/register", ctl.UserController.Register)
	r.POST("/login", ctl.UserController.Login)

	userRouter := r.Group("/user")
	{
		userRouter.GET("", ctl.UserController.CheckUser)
		userRouter.Use(mdl.UserMiddleware.Authentication())
		userRouter.Use(mdl.UserMiddleware.Authorization())
		userRouter.GET("/:userName", ctl.UserController.GetUser)
		userRouter.PUT("/:userName/ProfilPhoto", ctl.UserController.SetProfilePhoto)
		userRouter.PUT("/:userName", ctl.UserController.Update)
		userRouter.DELETE("/:userName", ctl.UserController.Delete)
	}
	donationRouter := r.Group("/donation")
	{
		donationRouter.Use(mdl.UserMiddleware.Authentication())
		donationRouter.POST("", ctl.DonationController.Create)
		donationRouter.GET("", ctl.DonationController.GetDonations)
		donationRouter.GET("/:donationId", ctl.DonationController.GetDonationById)
		donationRouter.Use(mdl.DonationMiddleware.Authorization())
		donationRouter.PUT("/:donationId", ctl.DonationController.EditDonation)
		donationRouter.DELETE("/:donationId", ctl.DonationController.DeleteDonation)

	}
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	return r.Run(":" + PORT)
}
