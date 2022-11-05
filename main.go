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
	userController := controllers.NewUserController(userRepository)

	controller := controllers.NewController(userController)

	err = routers.StartServer(controller)
	if err != nil {
		fmt.Println("error starting server", err)
		return
	}
}
