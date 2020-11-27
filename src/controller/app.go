package controller

import (
	"github.com/doter1995/user_center/src/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetApps(c *gin.Context) {
	var ps, p int
	s := service.App{}
	pageSize, _ := c.Params.Get("pageSize")
	page, _ := c.Params.Get("page")
	if c.ShouldBindJSON(&s) != nil {
		s = service.App{}
	}
	ps, err := strconv.Atoi(pageSize)
	if err != nil {
		ps = 1000
	}
	p, err = strconv.Atoi(page)
	if err != nil {
		p = 1
	}
	pages, err := s.FindApps(ps, p)
	if err != nil {
		c.JSON(500, err)
	}
	c.JSON(200, pages)
}

// CreateApp 创建APP
func CreateApp(c *gin.Context) {
	s := service.App{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(201, err)
	}
	status, err := s.CreateApp()
	c.JSON(200, Response{Code: status, Error: err.Error()})
}
