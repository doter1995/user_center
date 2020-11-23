package controller

import (
	"github.com/doter1995/user_center/src/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	s := service.UserLoginService{}
	if c.ShouldBindJSON(&s) == nil {
		c.JSON(http.StatusProxyAuthRequired, Response{Msg: "参数不合法"})
		return
	}
	code, _ := s.Login()
	if code < 0 {
		c.JSON(200, Response{Code: code, Msg: "登陆失败"})
	}
	c.JSON(200,ResponseWithToken{Code: code,Token: ""})
}
