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
	donationHistoryRepository := repository.NewDonationHistoryRepository(db)

	userController := controllers.NewUserController(userRepository)
	donationController := controllers.NewDonationController(donationRepository)
	donationRequestController := controllers.NewDonationRequestController(donationRequestRepository)
	donationHistoryController := controllers.NewDonationHistoryController(donationHistoryRepository)

	userMiddleware := middlewares.NewUserMiddleware(userRepository)
	donationMiddleware := middlewares.NewDonationMiddleware(donationRepository)
	donationRequestMiddleware := middlewares.NewDonationRequestMiddleware(donationRequestRepository)
	donationHistoryMiddleware := middlewares.NewDonationHistoryMiddleware(donationHistoryRepository)

	controller := controllers.NewController(userController, donationController, donationRequestController, donationHistoryController)
	middleware := middlewares.NewMiddleware(userMiddleware, donationMiddleware, donationRequestMiddleware, donationHistoryMiddleware)

	err, err2 := routers.StartServer(controller, middleware)
	if err != nil {
		fmt.Println("error starting server", err)
		return
	}
	if err2 != nil {
		fmt.Println("error starting server", err)
		return
	}
}
