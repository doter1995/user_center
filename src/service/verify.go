package service

import (
	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/model"
	"github.com/doter1995/user_center/src/tools"
)

func VerifyMD5(str string, code string) bool {
	slate := config.Config.Security.Token.Slate
	return tools.GetMD5Code(str, slate) == code
}
func NewToken(u model.User) (string, error) {
	token := config.Config.Security.Token
	c := tools.Claims{
		Username: u.Username,
		Status:   u.Status,
	}
	return tools.GenerateToken(c, token)
}
