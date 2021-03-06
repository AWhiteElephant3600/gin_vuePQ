package middleware

import (
	"gin_vuePQ/common"
	"gin_vuePQ/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization 头部
		tokenString := c.GetHeader("Authorization")

		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized,gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized,gin.H{
				"code": 401,
				"msg": "权限不足",
			})

			c.Abort()
			return
		}

		// 验证通过后获得claim中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userId)

		// 用户不存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized,gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}

		// 用户存在 将user的信息写入上下文
		c.Set("user",user)

		c.Next()

	}
}
