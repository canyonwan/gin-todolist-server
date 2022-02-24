package main

import (
	"canyonwan.top/gin_todolist_server/routes"
	"canyonwan.top/gin_todolist_server/setting"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.ForceConsoleColor()
	gin.SetMode(gin.DebugMode)
	// init viper
	initViperError := setting.InitViper()
	if initViperError != nil {
		return
	}

	// init mysql
	initMysql := setting.InitMysql()
	if initMysql != nil {
		return
	}
	defer setting.CloseDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	panic(r.Run())
}
