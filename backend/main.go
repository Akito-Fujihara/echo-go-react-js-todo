package main

import (
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/controller"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/db"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/repository"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/router"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.MewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
