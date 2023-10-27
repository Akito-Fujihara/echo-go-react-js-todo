package router

import (
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/controller"
	"github.com/labstack/echo/v4"
)

func MewRouter(uc controller.IUserController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	return e
}
