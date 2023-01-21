package routers

import (
	"dibagi/controllers"
	"dibagi/middlewares"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(ctl controllers.Controller, mdl middlewares.Middleware) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH"}
	config.AllowHeaders = []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Headers", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.GET("/", ctl.HomeController)
	// r.GET("/.well-known/pki-validation", ctl.HomeController)

	//manual login
	r.POST("/register", ctl.UserController.Register)
	r.POST("/login", ctl.UserController.Login)

	userRouter := r.Group("/user")
	{
		//check username or email exist
		userRouter.GET("", ctl.UserController.CheckUser)

		// Token authentication
		userRouter.Use(mdl.UserMiddleware.Authentication())
		// get one user data
		userRouter.GET("/:userName", ctl.UserController.GetUser)

		// check user have access or not
		userRouter.Use(mdl.UserMiddleware.Authorization())

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

		// send request to claim a donation from another user
		donationRouter.POST("/:donationId/request", mdl.DonationMiddleware.CheckDonator(), mdl.DonationRequestMiddleware.CheckIfExist(), ctl.DonationRequestController.Create)

		//get all request in donation
		donationRouter.GET("/:donationId/request", ctl.DonationRequestController.GetAllByDonationId)

		// get user submitted request
		donationRouter.GET("/request", ctl.DonationRequestController.GetAllByUserId)

		// get one donation request
		donationRouter.GET("/request/:donationRequestId", ctl.DonationRequestController.GetById)

		// cancel request
		donationRouter.DELETE("/request/:requestId", ctl.DonationRequestController.Delete)

		// donation authorization
		donationRouter.Use(mdl.DonationMiddleware.Authorization())
		// edit donation data
		donationRouter.PUT("/:donationId", ctl.DonationController.Edit)

		// delete donation
		donationRouter.DELETE("/:donationId", ctl.DonationController.Delete)

	}

	requestRouter := r.Group("/request")
	{
		requestRouter.Use(mdl.UserMiddleware.Authentication())
		// get all donation request
		requestRouter.GET("", ctl.DonationRequestController.GetAllByDonatorId)
		// confirm a request
		requestRouter.POST("/:donationRequestId", mdl.DonationRequestMiddleware.Authorization(), mdl.NotificationMiddleware.CheckIfExist(), ctl.DonationRequestController.Confirm)
		// reject a request
		requestRouter.DELETE("/:donationRequestId", mdl.DonationRequestMiddleware.Authorization(), mdl.NotificationMiddleware.CheckIfExist(), ctl.DonationRequestController.Reject)

	}
	notificationRouter := r.Group("/notification")
	{
		notificationRouter.Use(mdl.UserMiddleware.Authentication())
		// get notification
		notificationRouter.GET("", ctl.NotificationController.GetAllByUserId)
	}

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	go r.Run(":" + PORT)

	rTLS := gin.Default()

	rTLS.Use(cors.New(config))
	rTLS.GET("/", ctl.HomeController)
	// r.GET("/.well-known/pki-validation", ctl.HomeController)

	//manual login
	rTLS.POST("/register", ctl.UserController.Register)
	rTLS.POST("/login", ctl.UserController.Login)

	userRouter2 := rTLS.Group("/user")
	{
		//check username or email exist
		userRouter2.GET("", ctl.UserController.CheckUser)

		// Token authentication
		userRouter2.Use(mdl.UserMiddleware.Authentication())
		// get one user data
		userRouter2.GET("/:userName", ctl.UserController.GetUser)

		// check user have access or not
		userRouter2.Use(mdl.UserMiddleware.Authorization())

		//set user profil picture
		userRouter2.PUT("/:userName/ProfilPhoto", ctl.UserController.SetProfilePhoto)

		//edit user data
		userRouter2.PUT("/:userName", ctl.UserController.Update)

		//delete user account
		userRouter2.DELETE("/:userName", ctl.UserController.Delete)
	}
	donationRouter2 := rTLS.Group("/donation")
	{

		// Token authentication
		donationRouter2.Use(mdl.UserMiddleware.Authentication())

		// create a new donation
		donationRouter2.POST("", ctl.DonationController.Create)

		// get all donation
		donationRouter2.GET("", ctl.DonationController.GetAll)

		// get one donation by id
		donationRouter2.GET("/:donationId", ctl.DonationController.GetDonationById)

		// send request to claim a donation from another user
		donationRouter2.POST("/:donationId/request", mdl.DonationMiddleware.CheckDonator(), mdl.DonationRequestMiddleware.CheckIfExist(), ctl.DonationRequestController.Create)

		//get all request in donation
		donationRouter2.GET("/:donationId/request", ctl.DonationRequestController.GetAllByDonationId)

		// get user submitted request
		donationRouter2.GET("/request", ctl.DonationRequestController.GetAllByUserId)

		// get one donation request
		donationRouter2.GET("/request/:donationRequestId", ctl.DonationRequestController.GetById)

		// cancel request
		donationRouter2.DELETE("/request/:requestId", ctl.DonationRequestController.Delete)

		// donation authorization
		donationRouter2.Use(mdl.DonationMiddleware.Authorization())
		// edit donation data
		donationRouter2.PUT("/:donationId", ctl.DonationController.Edit)

		// delete donation
		donationRouter2.DELETE("/:donationId", ctl.DonationController.Delete)

	}

	requestRouter2 := r.Group("/request")
	{
		requestRouter2.Use(mdl.UserMiddleware.Authentication())
		// get all donation request
		requestRouter2.GET("", ctl.DonationRequestController.GetAllByDonatorId)
		// confirm a request
		requestRouter2.POST("/:donationRequestId", mdl.DonationRequestMiddleware.Authorization(), mdl.NotificationMiddleware.CheckIfExist(), ctl.DonationRequestController.Confirm)
		// reject a request
		requestRouter2.DELETE("/:donationRequestId", mdl.DonationRequestMiddleware.Authorization(), mdl.NotificationMiddleware.CheckIfExist(), ctl.DonationRequestController.Reject)
	}
	notificationRouter2 := rTLS.Group("/notification")
	{
		notificationRouter2.Use(mdl.UserMiddleware.Authentication())
		// get notification
		notificationRouter2.GET("", ctl.NotificationController.GetAllByUserId)
	}

	// return
	return rTLS.RunTLS(":443", "./cert/certificate.crt", "./cert/private.key")

}
