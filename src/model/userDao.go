package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;unique_index"`
	Password string `gorm:"type:varchar(160);not null"`
	Email    string `gorm:"type:varchar(160)"`
	Icon     string
	Info     string
	AuthCode string `gorm:"type:varchar(100);not null"`
}

//FindUserByUsername 通过username查找User
func FindUserByUsername(username string) User {
	var u User
	DB.Where("username = ?", username).First(&u)
	return u
}

func FindUserById(id uint) User {
	var u User
	DB.Where("id = ?", id).First(&u)
	return u
}

//CreateUser 创建用户
func CreateUser(u User) User {
	var u1 User
	DB.First(&u1, User{Username: u.Username})
	if u1.Username == "" {
		DB.Create(&u)
		return u
	}
	return u1
}
