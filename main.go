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
	donationRequestRepository := repository.NewDonationRequestRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)

	userController := controllers.NewUserController(userRepository)
	donationController := controllers.NewDonationController(donationRepository)
	donationRequestController := controllers.NewDonationRequestController(donationRequestRepository)
	notificationController := controllers.NewNotificationController(notificationRepository)

	userMiddleware := middlewares.NewUserMiddleware(userRepository)
	donationMiddleware := middlewares.NewDonationMiddleware(donationRepository)
	donationRequestMiddleware := middlewares.NewDonationRequestMiddleware(donationRequestRepository)
	notificationMiddleware := middlewares.NewNotificationMiddleware(notificationRepository)

	controller := controllers.NewController(userController, donationController, donationRequestController, notificationController)
	middleware := middlewares.NewMiddleware(userMiddleware, donationMiddleware, donationRequestMiddleware, notificationMiddleware)

	err = routers.StartServer(controller, middleware)
	if err != nil {
		fmt.Println("error starting server", err)
		return
	}
}
