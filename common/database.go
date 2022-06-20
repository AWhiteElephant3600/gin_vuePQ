package common

import (
	"fmt"
	"gin_vuePQ/model"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
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

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("failed to connect database,err="+err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
