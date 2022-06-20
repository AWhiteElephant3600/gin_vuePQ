package controller

import (
	"gin_vuePQ/model"
	"gin_vuePQ/repository"
	"gin_vuePQ/response"
	"gin_vuePQ/vo"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx,nil,"数据验证错误,分类名称必填")
		return
	}
	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}

	response.Success(ctx,gin.H{"category": category},"添加成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest

	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx,nil,"数据验证错误,分类名称必填")
		return
	}

	categoryId, err := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx,nil,"分类不存在")
		return
	}

	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, gin.H{"category": category},"修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, nil, "分类不存在")
		return
	}

	response.Success(ctx,gin.H{"category": category},"")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, err := strconv.Atoi(ctx.Params.ByName("id"))

	_, err = c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx,nil,"分类不存在")
		return
	}

	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx,nil,"删除失败")
		return
	}

	response.Success(ctx,nil,"删除成功")
}

