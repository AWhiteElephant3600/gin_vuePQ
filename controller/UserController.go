package controller

import (
	"gin_vuePQ/common"
	"gin_vuePQ/dto"
	"gin_vuePQ/model"
	"gin_vuePQ/response"
	"gin_vuePQ/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	c.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)

	if isTelephoneExist(DB, telephone) {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户已存在")
		return
	}

	// 创建用户
	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"加密错误")
		return
	}

	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hasedPassword),
	}

	DB.Create(&newUser)

	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"发放token失败")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	response.Success(c, gin.H{"token": token},"注册成功")

}

func Login(c *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	c.Bind(&requestUser)

	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); err != nil {
		response.Fail(c,nil,"密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c,http.StatusInternalServerError,500,nil,"系统异常")
		log.Printf("token generate error: %v",err)
		return
	}

	// 返回结果
	response.Success(c,gin.H{"token": token},"登陆成功")

}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK,gin.H{"code": 200,"data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB,telephone string) bool {
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}