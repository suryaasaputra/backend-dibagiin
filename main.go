package main

import (
	"dibagi/controllers"
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

	controller := controllers.NewController(userController, donationController)

	err = routers.StartServer(controller)
	if err != nil {
		fmt.Println("error starting server", err)
		return
	}
}
