package router

import (
	"github.com/doter1995/user_center/src/controller"
	"github.com/doter1995/user_center/src/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(s *gin.Engine) {
	api := s.Group("api")
	initUserRouter(api)
	api.Use(middleware.Authentication)
	initUserInfoRouter(api)
	initAppInfoRouter(api)
	api.Use(middleware.NextMiddleware)
}

func initUserRouter(router *gin.RouterGroup) {
	router.POST("/register", controller.UserRegister) // 用户注册
	router.POST("/login", controller.UserLogin)       // 登录
}

func initUserInfoRouter(router *gin.RouterGroup) {
	router.POST("/changePassword")
	router.GET("/user/:id", controller.GetUserById)
	router.PUT("/user/:id")
	router.DELETE("/user/:id")
	router.GET("/users/:pageSize/:page", controller.GetUsers)
}
func initAppInfoRouter(router *gin.RouterGroup) {
	router.GET("/app/:id")
	router.POST("/app/:id", controller.CreateApp)
	//router.PUT("/app/:id", controller.UpdateApp)
	//router.DELETE("/app/:id", controller.DeleteApp)
	router.GET("/apps/:pageSize/:page", controller.GetApps)
}
