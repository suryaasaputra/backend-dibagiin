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
		//check username or email exist
		userRouter.GET("", ctl.UserController.CheckUser)

		// Token authentication
		userRouter.Use(mdl.UserMiddleware.Authentication())

		// check user have access or not
		userRouter.Use(mdl.UserMiddleware.Authorization())

		// get one user data
		userRouter.GET("/:userName", ctl.UserController.GetUser)

		//set user profil picture
		userRouter.PUT("/:userName/ProfilPhoto", ctl.UserController.SetProfilePhoto)

		//edit user data
		userRouter.PUT("/:userName", ctl.UserController.Update)

		//delete user account
		userRouter.DELETE("/:userName", ctl.UserController.Delete)
	}
	donationRouter := r.Group("/donation")
	{

		// Token authentication
		donationRouter.Use(mdl.UserMiddleware.Authentication())

		// create a new donation
		donationRouter.POST("", ctl.DonationController.Create)

		// get all donation
		donationRouter.GET("", ctl.DonationController.GetAll)

		// get one donation by id
		donationRouter.GET("/:donationId", ctl.DonationController.GetDonationById)

		//get all request in donation
		donationRouter.GET("/:donationId/request", ctl.DonationController.GetDonationById)

		// request to claim a donation from another user
		donationRouter.POST("/request/:donationId", mdl.DonationRequestMiddleware.CheckIfExist(), ctl.DonationRequestController.Create)

		// get all donation request
		donationRouter.GET("/request", ctl.DonationRequestController.GetAllByDonatorId)

		// get one donation request
		donationRouter.GET("/request/:donationRequestId", ctl.DonationRequestController.GetById)

		// confirm a request
		donationRouter.PUT("/request/:donationRequestId", ctl.DonationRequestController.Confirm)

		donationRouter.Use(mdl.DonationMiddleware.Authorization())
		// edit donation data
		donationRouter.PUT("/:donationId", ctl.DonationController.Edit)

		// delete donation
		donationRouter.DELETE("/:donationId", ctl.DonationController.Delete)

	}

	requestRouter := r.Group("/request")
	{
		// get user donation request

		requestRouter.GET("/:userId", ctl.DonationRequestController.GetAllByUserId)

	}

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	return r.Run(":" + PORT)
}
