package common

import (
	"fmt"
	"gvp/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB{
	username:=viper.GetString("datasource.username")
  	password:=viper.GetString("datasource.password")
	host:=viper.GetString("datasource.host")
  	port:=viper.GetString("datasource.port")
  	database:=viper.GetString("datasource.database")
  	charset:=viper.GetString("datasource.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
username,password,host,port,database,charset)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic("failed to connect to database,error: "+err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB=db
	return db
}

func GetDB() *gorm.DB{
	return DB
}