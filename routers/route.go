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
	r.StaticFile("/.well-known/pki-validation/67947117BCCC6E7F5087A16C09F9B136.txt", "./67947117BCCC6E7F5087A16C09F9B136.txt")
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

		// request to claim a donation from another user
		donationRouter.POST("/:donationId/request", mdl.DonationMiddleware.CheckDonator(), mdl.DonationRequestMiddleware.CheckIfExist(), ctl.DonationRequestController.Create)

		//get all request in donation
		donationRouter.GET("/:donationId/request", ctl.DonationRequestController.GetAllByDonationId)

		// get all donation request
		donationRouter.GET("/request", ctl.DonationRequestController.GetAllByDonatorId)

		// get one donation request
		donationRouter.GET("/request/:donationRequestId", ctl.DonationRequestController.GetById)

		// confirm a request
		donationRouter.POST("/request/:donationRequestId/confirm", mdl.DonationRequestMiddleware.Authorization(), mdl.DonationHistoryMiddleware.CheckIfExist(), ctl.DonationRequestController.Confirm)
		// reject a request
		donationRouter.POST("/request/:donationRequestId/reject", mdl.DonationRequestMiddleware.Authorization(), mdl.DonationHistoryMiddleware.CheckIfExist(), ctl.DonationRequestController.Reject)

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
		// get user submitted request
		requestRouter.GET("", ctl.DonationRequestController.GetAllByUserId)
		requestRouter.DELETE("/:requestId", ctl.DonationRequestController.Delete)
	}
	historyRouter := r.Group("/history")
	{
		historyRouter.Use(mdl.UserMiddleware.Authentication())
		// get history
		historyRouter.GET("", ctl.DonationHistoryController.GetAllByUserId)
	}

	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// return r.Run(":" + PORT)
	return r.RunTLS(":"+PORT, "./cert/certificate.crt", "./cert/private.key")

}
