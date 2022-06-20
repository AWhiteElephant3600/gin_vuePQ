package middleware

import (
	"fmt"
	"gin_vuePQ/response"
	"github.com/gin-gonic/gin"
)


func RecoveryMiddleware() gin.HandlerFunc {
	// 全局错误捕获
	return func(ctx *gin.Context) {
		defer func() { // 最后处理 defer
			if err := recover(); err != nil {
				response.Fail(ctx,nil,fmt.Sprint(err))
			}
		}()

		ctx.Next()
	}
}
