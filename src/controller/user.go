package controller

import (
	"github.com/doter1995/user_center/src/service"
	"github.com/gin-gonic/gin"
	"strconv"
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
func UserRegister(c *gin.Context) {
	s := service.UserRegisterService{}
	if c.ShouldBindJSON(&s) != nil {
		c.JSON(200, Response{Code: -1, Msg: "用户存在"})
		return
	}
	s.Register()
	c.JSON(200, Response{Code: 0})
	return
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := service.FindUserById(id)
	if id == "" || err != nil {
		c.JSON(403, nil)
		return
	}
	c.JSON(200, user)
}
func GetUsers(c *gin.Context) {
	var ps, p int
	s := service.User{}
	pageSize, _ := c.Params.Get("pageSize")
	page, _ := c.Params.Get("page")
	if c.ShouldBindJSON(&s) != nil {
		s = service.User{}
	}
	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		ps = 1000
	}
	p, err = strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	pages, err := s.FindUsers(ps, p)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, pages)

}
