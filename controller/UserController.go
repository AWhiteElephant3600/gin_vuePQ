package controller

import (
	"gin_vuePQ/model"
	"gin_vuePQ/response"
	"gin_vuePQ/service"
	"gin_vuePQ/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type IUserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Info(ctx *gin.Context)
}

type UserController struct {
	Service service.IUserService
}

func NewUserController() IUserController {
	service := service.NewUserService()
	return UserController{Service: service}
}

// UserControllerRegister 用户注册列表接口
// @Summary 用户注册接口
// @Description 可用户注册
// @Tags 可用户注册
// @Accept application/json
// @Produce application/json
// @Param object query model.User false "查询参数"
// @Security ApiKeyAuth
// @Success 200
// @Router /api/auth/register [post]
func (u UserController) Register(ctx *gin.Context) {
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	if len(name) == 0 {
		requestUser.Name = util.RandomString(10)
	}

	log.Println(name, telephone, password)

	//if isTelephoneExist(u.Repository.DB, telephone) {
	//	response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户已存在")
	//	return
	//}

	token, err := u.Service.Register(requestUser)
	if err != nil {
		response.Fail(ctx, gin.H{"err": err.Error()}, "注册失败")
		return

	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

func (u UserController) Login(ctx *gin.Context) {
	// 接收前端传过来的手机号密码
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	telephone := requestUser.Telephone

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	// 将登录的user实体传到持久层的Login方法中
	token, err := u.Service.Login(requestUser)
	if err != nil {
		response.Fail(ctx, gin.H{"err": err.Error()}, "登陆失败")
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

func (u UserController) Info(ctx *gin.Context) {
	userDto := u.Service.Info(ctx)
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": userDto}})
}

/*func (u UserController) Register(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password

	if len(telephone) != 11 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"密码不能少于6位")
	}

	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)

	if isTelephoneExist(DB, telephone) {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户已存在")
		return
	}

	// 创建用户
	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx,http.StatusInternalServerError,500,nil,"加密错误")
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
		response.Response(ctx,http.StatusInternalServerError,500,nil,"发放token失败")
		log.Printf("token generate error: %v", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token": token},"注册成功")

}

func (u UserController) Login(ctx *gin.Context) {
	DB := common.GetDB()
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	telephone := requestUser.Telephone
	password := requestUser.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"手机号必须为11位")
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx,http.StatusUnprocessableEntity,422,nil,"用户不存在")
		return
	}

	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password)); err != nil {
		response.Fail(ctx,nil,"密码错误")
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx,http.StatusInternalServerError,500,nil,"系统异常")
		log.Printf("token generate error: %v",err)
		return
	}

	// 返回结果
	response.Success(ctx,gin.H{"token": token},"登陆成功")
}

func (u UserController) Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK,gin.H{"code": 200,"data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}*/

/*func isTelephoneExist(db *gorm.DB,telephone string) bool {
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}*/
