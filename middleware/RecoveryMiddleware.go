package middleware

import (
	"fmt"
	"gin_vuePQ/response"
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	// 全局错误捕获
	return func(ctx *gin.Context) {
		defer func() { // 最后处理 (defer)
			// recover()捕获，并将允许用if err != nil错误信息打印出来
			if err := recover(); err != nil {
				response.Fail(ctx, nil, fmt.Sprint(err))
			}
		}()

		// 后面的回调函数继续进行
		ctx.Next()
	}
}
