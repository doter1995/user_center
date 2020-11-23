package controller

import (
	"github.com/doter1995/user_center/src/service"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	s := service.UserLoginService{}
	if c.ShouldBindJSON(&s) != nil {
		c.JSON(201, Response{Msg: "参数不合法"})
		return
	}
	code, token := s.Login()
	if code < 0 {
		c.JSON(200, Response{Code: code, Msg: "登陆失败"})
		return
	}
	c.JSON(200, ResponseWithToken{Code: code, Token: token})
}
