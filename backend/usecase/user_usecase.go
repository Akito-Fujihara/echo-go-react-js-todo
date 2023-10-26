package usecase

import (
	"os"
	"time"

	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/model"
	"github.com/Akito-Fujihara/echo-go-react-js-todo/backend/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUseCase interface {
	SignUp(user *model.User) (model.UserResponse, error)
	Login(user *model.User) (model.UserResponse, error)
}

type userUseCase struct {
	ur repository.IUserRepository
}

func NewUserUseCase(ur repository.IUserRepository) IUserUseCase {
	return &userUseCase{ur}
}

func (uu *userUseCase) SignUp(user *model.User) (model.UserResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PassWord), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{
		Email:    user.Email,
		PassWord: string(hash),
	}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	return model.UserResponse{}, nil
}
func (uu *userUseCase) Login(user model.User) (string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.PassWord), []byte(user.PassWord))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
