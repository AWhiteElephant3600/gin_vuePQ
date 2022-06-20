package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 定义全局的DB变量
var DB *gorm.DB

func InitDB() *gorm.DB {
	// 从配置文件中读取相对应的数据库连接所需的变量值--postgresql
	//driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	user := viper.GetString("datasource.user")
	password := viper.GetString("datasource.password")
	dbname := viper.GetString("datasource.dbname")
	port := viper.GetString("datasource.port")
	sslmode := viper.GetString("datasource.sslmode")
	timezone := viper.GetString("datasource.TimeZone")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
		timezone)
	// 连接,得到连接对象
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database,err=" + err.Error())
	}

	// 将连接对象赋值给DB
	DB = db
	return db
}

/*
提供一个对外的取得DB连接对象的方法
*/
func GetDB() *gorm.DB {
	return DB
}
