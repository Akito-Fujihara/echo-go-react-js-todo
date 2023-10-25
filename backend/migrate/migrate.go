package main

import (
	"fmt"

	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/db"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("db migrate success")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
