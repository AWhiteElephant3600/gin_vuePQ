package main

import (
	"gin_vuePQ/common"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// @title cbx
// @version 1.0
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name cbx
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	// 初始化配置
	InitConfig()
	// 初始化DB连接数据库对象
	common.InitDB()

	// 初始化路由
	r := gin.Default()

	// 进行路由分组，路由配置
	r = CollectRoute(r)

	// 读取项目端口号的配置
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())

}

func InitConfig() {
	// 获得项目根路径
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
