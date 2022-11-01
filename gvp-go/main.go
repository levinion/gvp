package main

import (
	"gvp/common"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	panic(r.Run(":" + port))
}

func initConfig() {
	dir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(dir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
