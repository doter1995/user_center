package main

import (
	"path/filepath"

	"github.com/doter1995/user_center/src/model"

	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/router"
	"github.com/doter1995/user_center/src/tools"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig(filepath.Base("./config.yaml"))
	tools.InitLog()
	model.InitDB()
	startService()
}
func startService() {
	s := gin.Default()
	s.Use(gzip.Gzip(gzip.DefaultCompression))
	router.InitRouter(s)
	s.Run(":" + config.Config.Server.Port)
}
