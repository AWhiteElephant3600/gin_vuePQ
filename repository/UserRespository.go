package repository

import (
	"errors"
	"gin_vuePQ/common"
	"gin_vuePQ/dto"
	"gin_vuePQ/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() UserRepository {
	return UserRepository{DB: common.GetDB()}
}

func (u UserRepository) Register(user model.User) (string, error) {

	telephone := user.Telephone
	password := user.Password
	name := user.Name

	if u.isTelephoneExist(telephone) {
		return "", errors.New("用户已存在")
	}

	// 创建用户
	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("加密失败")
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}

	u.DB.Create(&newUser)

	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		log.Printf("token generate error: %v", err)
		return "", errors.New("发放token失败")
	}

	// 返回结果
	return token, nil
}

func (u UserRepository) Login(user model.User) (string, error) {

	telephone := user.Telephone
	password := user.Password

	u.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		return "", errors.New("用户不存在")
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("密码错误")
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		log.Printf("token generate error: %v", err)
		return "", errors.New("系统异常")
	}

	// 返回结果
	return token, nil
}

func (u UserRepository) Info(ctx *gin.Context) dto.UserDto {
	user, _ := ctx.Get("user")
	return dto.ToUserDto(user.(model.User))
}

func (u UserRepository) isTelephoneExist(telephone string) bool {
	var user model.User
	u.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
