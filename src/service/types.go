package service

// UserLoginService 用户登陆字段校验
type UserLoginService struct {
	Username string
	Password string
	AuthCode string //OTP 生成的code
}
