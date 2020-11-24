package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;unique_index"`
	Password string `gorm:"type:varchar(160);not null"`
	Email    string `gorm:"type:varchar(160)"`
	Status   int    `gorm:"type:int(2)`
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

func FindUserById(id string) (User, error) {
	var u User
	result := DB.First(&u, id)
	return u, result.Error
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

func FindUser(pageSize int, page int, u User) (UserPagination, error) {
	users := make([]User, 0)
	userPage := UserPagination{}
	var total int64 = 0
	db := DB.Model(&User{})
	if u.ID != 0 {
		db.Where("ID= ?", u.ID)
	}
	if u.Username != "" {
		db.Where("username = ?", u.Username)
	}

	if u.Status != 0 {
		db.Where("status = ?", u.Status)
	}
	db.Count(&total)
	if pageSize > 0 {
		userPage.PageSize = pageSize
		db.Limit(pageSize)
	} else {
		userPage.PageSize = 1000
		db.Limit(1000)
	}
	if page > 1 {
		userPage.Page = page
		db.Offset((page - 1) * pageSize)
	} else {
		userPage.Page = 1
		db.Offset(0)
	}

	if err := DB.Find(&users).Error; err != nil {
		return userPage, err
	}
	userPage.Data = users
	userPage.Total = int(total)
	return userPage, nil
}
