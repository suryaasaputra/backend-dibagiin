package routers

import (
	"dibagi/controllers"
	"dibagi/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func StartServer(ctl controllers.Controller) error {

	r := gin.Default()
	r.LoadHTMLGlob("*.html")
	r.GET("/", ctl.HomeController)

	//manual login
	r.POST("/register", ctl.UserController.Register)
	r.POST("/login", ctl.UserController.Login)

	userRouter := r.Group("/user")
	{
		userRouter.GET("/:userName", ctl.UserController.GetUser)
		userRouter.GET("", ctl.UserController.CheckUser)
		userRouter.Use(middlewares.Authentication())
		userRouter.Use(middlewares.UserAuthorization())
		userRouter.PATCH("/:userName/UpdateProfilPhoto", ctl.UserController.SetProfilePhoto)
		userRouter.PUT("/:userName", ctl.UserController.Update)
		userRouter.DELETE("/:userName", ctl.UserController.Delete)
	}
	donationRouter := r.Group("/donation")
	{
		donationRouter.POST("", middlewares.Authentication(), ctl.DonationController.Create)
		donationRouter.GET("", middlewares.Authentication(), ctl.DonationController.GetDonations)

	}
	var PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	return r.Run(":" + PORT)
}
