package user_center

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	startService()
}
func startService() {
	s := gin.Default()
	s.Use(gzip.Gzip(gzip.DefaultCompression))
	InitRouter(s)
}
