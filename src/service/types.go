package service

// UserLoginService 用户登陆字段校验
type UserLoginService struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthCode string `json:"authCode" binding:"required,min=6,max=6,number` //OTP 生成的code
}

// UserRegisterService 用户登陆字段校验
type UserRegisterService struct {
	Username string `json:"username" binding:"required,min=6,max=16"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

//User 用于User查询
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   int    `json:"status"` //代表状态 0 正常，1代表锁定
}

//App 用与应用管理查询
type App struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"type:varchar(100);not null;unique_index"`
	Info   string
	Status int `gorm:"type:int(2)` //代表状态 0 正常，1代表锁定
	Path   string
	Code   string `gorm:"type:varchar(100);not null" json:"-"` //
}

