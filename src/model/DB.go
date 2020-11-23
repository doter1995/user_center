package model

import (
	"github.com/doter1995/user_center/src/config"
	"github.com/doter1995/user_center/src/tools"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB 数据库连接
var DB *gorm.DB

//InitDB 初始化数据库连接
func InitDB() {
	c := config.Config
	var (
		db  *gorm.DB
		err error
	)
	switch c.Database.Type {
	case "UNSET", "sqlite", "sqlite3":
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(c.Database.Url), &gorm.Config{}) //todo:process user

	default:
		tools.Logger.Error().Str("Database.Type", c.Database.Type).Msg("暂不支持数据库类型")
	}
	if err != nil {
		tools.Logger.Err(err).Msg("连接数据库不成功")
	}
	tools.Logger.Info().Msg("成功连接到数据库")
	DB = db

	migrateDB()
}

func migrateDB() {
	slate := config.Config.Security.Token.Slate
	DB.AutoMigrate(&User{})
	u := CreateUser(User{Username: "admin", Password: tools.GetMD5Code("doter1995", slate), AuthCode: "MPNP3K3LBG4EUGDZMZUOIODSKR2PXBM6"})
	if u.Username == "" {
		tools.Logger.Info().Msg("成功创建admin用户")
	} else {
		tools.Logger.Info().Msg("admin已存在")
	}
	tools.Logger.Info().Msg("自动迁移数据库成功！")
}
