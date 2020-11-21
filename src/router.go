package user_center

import (
	"github.com/doter1995/user_center/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(s *gin.Engine)  {
	api := s.Group("api")
	initUserRouter(api)
	initVerifyRouter(api)
	api.Use(middleware.Authentication)
	api.Use(middleware.NextMiddleware)
}

func initUserRouter(router *gin.RouterGroup)  {
	router.POST("/register") // 用户注册
	router.POST("/login") // 登录
}
func initVerifyRouter(router *gin.RouterGroup)  {
	router.POST("/oneTimeCode") // 一次验证码
}
