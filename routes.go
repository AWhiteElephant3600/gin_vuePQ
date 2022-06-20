package main

import (
	"gin_vuePQ/controller"
	"gin_vuePQ/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	// 设置全局中间件,跨域处理中间件，全局err捕获处理中间件
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	userController := controller.NewUserController()
	r.POST("/api/auth/register", userController.Register)
	r.POST("/api/auth/login", userController.Login)
	// AuthMiddleware认证中间件
	r.GET("/api/auth/info", middleware.AuthMiddleware(), userController.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update)
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.DELETE("/:id", categoryController.Delete)

	postRoutes := r.Group("/post")
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("/:id", postController.Update)
	postRoutes.GET("/:id", postController.Show)
	postRoutes.DELETE("/:id", postController.Delete)
	postRoutes.GET("/page/list", postController.PageList)

	return r
}
