package service

import (
	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/tools"
)

func VerifyMD5(str string, code string) bool {
	slate := config.Config.Security.Token.Slate
	return tools.GetMD5Code(str, slate) == code
}
