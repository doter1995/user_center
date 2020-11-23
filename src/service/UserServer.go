package service

import (
	"github.com/doter1995/user_center/src/model"
	"github.com/doter1995/user_center/src/tools"
)

func (s *UserLoginService) Login() (int, error) {
	u := model.FindUserByUsername(s.Username)
	if u.Username != s.Username {
		return -1, nil // 用户不存在
	}
	if !VerifyMD5(s.Password, u.Password) {
		return -2, nil // 密码错误
	}
	if !tools.VerifyOTPCode(u.AuthCode, s.AuthCode) {
		return -3, nil //验证码错误
	}
	return 0, nil
}
