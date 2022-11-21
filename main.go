package main

import (
	"dibagi/controllers"
	"dibagi/middlewares"
	"dibagi/repository"
	"dibagi/routers"
	"fmt"
)

func main() {
	db, err := repository.NewDB()
	if err != nil {
		fmt.Println("error starting database", err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	donationRepository := repository.NewDonationRepository(db)

	userController := controllers.NewUserController(userRepository)
	donationController := controllers.NewDonationController(donationRepository)

	userMiddleware := middlewares.NewUserMiddleware(userRepository)
	donationMiddleware := middlewares.NewDonationMiddleware(donationRepository)

	controller := controllers.NewController(userController, donationController)
	middleware := middlewares.NewMiddleware(userMiddleware, donationMiddleware)

	err = routers.StartServer(controller, middleware)
	if err != nil {
		fmt.Println("error starting server", err)
		return
	}
}
