package service

// UserLoginService 用户登陆字段校验
type UserLoginService struct {
	Username string  `json:"username"`
	Password string	`json:"password"`
	AuthCode string `json:"authCode" binding:"required,min=6,max=6,number`//OTP 生成的code
}
