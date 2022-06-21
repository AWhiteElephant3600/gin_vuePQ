package service

import (
	"gin_vuePQ/dto"
	"gin_vuePQ/model"
	"gin_vuePQ/repository"
	"github.com/gin-gonic/gin"
)

type IUserService interface {
	Register(user model.User) (string, error)
	Login(user model.User) (string, error)
	Info(ctx *gin.Context) dto.UserDto
}

type UserService struct {
	Repository repository.UserRepository
}

func NewUserService() IUserService {
	userRepository := repository.NewUserRepository()
	userRepository.DB.AutoMigrate(model.User{})
	return UserService{Repository: userRepository}
}

func (u UserService) Register(user model.User) (string, error) {
	token, err := u.Repository.Register(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u UserService) Login(user model.User) (string, error) {
	token, err := u.Repository.Login(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u UserService) Info(ctx *gin.Context) dto.UserDto {
	userInfo := u.Repository.Info(ctx)
	return userInfo
}
