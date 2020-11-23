package middleware

import (
	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"strings"
)

func Authentication(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "请求头中auth为空",
		})
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 0,
			"msg":  "请求头中auth格式有误",
		})
		c.Abort()
		return
	}
	_, err := tools.VerifyToken(parts[1], config.Config.Security.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "1"})
	}

	c.Next()
}

//NextMiddleware 向下进行代理代理
func NextMiddleware(c *gin.Context) {
	Host := "127.0.0.1"
	proxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = Host
			req.Host = Host
		},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
