package middleware

import (
	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/model"
	"github.com/doter1995/user_center/src/service"
	"github.com/doter1995/user_center/src/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
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
	token := parts[1]
	claims, err := tools.VerifyToken(token, config.Config.Security.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "1"})
		c.Abort()
		return
	}
	//以上为token校验
	//以下为redis对token校验
	tokenStr, err := service.RedisGetToken(claims.Username)
	if err != nil || tokenStr != token {
		c.JSON(http.StatusUnauthorized, gin.H{"code": "1"})
		c.Abort()
		return
	}
	//判断token是否需要刷新 暂定1.5-2小时内会出发token更新
	if claims.IssuedAt+int64(time.Minute*90) >= time.Now().Unix() {
		tokenStr, err := service.NewToken(model.User{
			Username: claims.Username,
			Status:   claims.Status,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": "3", "msg": "token更新失败"})
			c.Abort()
			return
		}
		c.Header("token", tokenStr)
	} else { // redis token 过期了
		c.JSON(http.StatusUnauthorized, gin.H{"code": "1", "msg": "token已经过期"})
		c.Abort()
		return
	}
	c.Set("tokenName", claims.Username)
	c.Set("userName", claims.Username)
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
