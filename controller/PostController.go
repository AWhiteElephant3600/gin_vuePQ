package controller

import (
	"gin_vuePQ/model"
	"gin_vuePQ/repository"
	"gin_vuePQ/response"
	"gin_vuePQ/vo"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(ctx *gin.Context)
}

type PostController struct {
	Repository repository.PostRepository
}

func NewPostController() IPostController {
	repository := repository.NewPostRepository()
	// 绑定User结构体,根据结构体属性的声明自动生成对应的表
	repository.DB.AutoMigrate(model.Post{})
	return PostController{Repository: repository}
}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取登陆用户
	user, _ := ctx.Get("user")

	// 创建post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	if err := p.Repository.Create(post); err != nil {
		panic(err)
		return
	}

	/*	if err := p.Repository.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}*/

	response.Success(ctx, nil, "创建成功")

}

func (p PostController) Update(ctx *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, nil, "数据验证错误")
		return
	}

	// 获取path中的id
	postId := ctx.Params.ByName("id")

	var post model.Post
	if err := p.Repository.DB.Preload("Category").Where("id = ?", postId).First(&post).Error; err != nil {
		response.Fail(ctx, nil, "文章不存在")
		return
	}

	// 判断当前用户是否为文章的作者
	// 获取登录用户
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "文章不属于你,请勿非法操作")
		return
	}

	post.Content = requestPost.Content
	post.Title = requestPost.Title
	post.HeadImg = requestPost.HeadImg
	post.CategoryId = requestPost.CategoryId

	if err := p.Repository.Update(post); err != nil {
		response.Fail(ctx, gin.H{"err": err.Error()}, "更新失败")
		return
	}
	// 更新文章
	/*	if err := p.Repository.DB.Save(&post).Error; err != nil {
		response.Fail(ctx,nil,"更新失败")
		return
	}*/

	response.Success(ctx, gin.H{"post": post}, "更新成功")

}

func (p PostController) Show(ctx *gin.Context) {
	// 获取path中的id
	postId := ctx.Params.ByName("id")

	post, err := p.Repository.Show(postId)
	if err != nil {
		response.Fail(ctx, gin.H{"err": err.Error()}, "文章不存在")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "成功")
}

func (p PostController) Delete(ctx *gin.Context) {
	postId := ctx.Params.ByName("id")

	post, err := p.Repository.Show(postId)
	if err != nil {
		response.Fail(ctx, gin.H{"err": err.Error()}, "文章不存在")
		return
	}
	// 判断当前用户是否文章的作者
	// 获取登录用户
	user, _ := ctx.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(ctx, nil, "文章不属于你,请勿非法操作")
		return
	}

	err = p.Repository.Delete(post)
	if err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, gin.H{"post": post}, "删除成功")

}

func (p PostController) PageList(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	posts, total, err := p.Repository.PageList(pageNum, pageSize)
	if err != nil {
		response.Fail(ctx, nil, "分页展示失败")
	}

	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")
}
