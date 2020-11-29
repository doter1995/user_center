package tools

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger 全局日志
var Logger zerolog.Logger

//InitLog 初始化log
func InitLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//f, _ := os.Create("log.log")
	logger :=  log.Output(os.Stdout)
	Logger = logger
}
