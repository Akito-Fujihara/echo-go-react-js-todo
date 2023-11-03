package main

import (
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/controller"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/db"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/repository"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/router"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/usecase"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/validator"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
