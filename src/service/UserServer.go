package service

import (
	"github.com/doter1995/user_center/src/model"
	"github.com/doter1995/user_center/src/tools"
)

func (s *UserLoginService) Login() (int, string) {
	u := model.FindUserByUsername(s.Username)
	if u.Username != s.Username {
		return -1, "" // 用户不存在
	}
	if !VerifyMD5(s.Password, u.Password) {
		return -2, "" // 密码错误
	}
	if !tools.VerifyOTPCode(u.AuthCode, s.AuthCode) {
		return -3, "" //验证码错误
	}
	token, err := NewToken(u)
	if err != nil {
		return -4, ""
	}
	return 0, token
}
