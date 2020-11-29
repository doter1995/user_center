package model

import "gorm.io/gorm"

type App struct {
	gorm.Model `json:"-"`
	ID         uint   `gorm:"primarykey" json:"id"`
	Name       string `gorm:"type:varchar(100);not null;unique_index" json:"name"`
	Info       string `json:"info"`
	Path       string `json:"path"`
	Status     int    `gorm:"type:int(2)" json:"status"`              //代表状态 0 正常，1代表锁定
	Code       string `gorm:"type:varchar(100);not null" json:"code"` //预留字段 暂未想好后续实现
}

func CreateApp(app App) App {
	if FindApp(app) {
		DB.Create(&app)
		return app
	}
	return App{}
}

func FindApp(app App) bool {
	var a1 App
	DB.First(&a1, app)
	return a1.Name == app.Name
}

//FindApp 查找App
func FindApps(pageSize int, page int, app App) (AppPagination, error) {
	apps := make([]App, 0)
	appPage := AppPagination{}
	var total int64 = 0
	db := DB.Model(&App{})
	if app.ID != 0 {
		db.Where("ID= ?", app.ID)
	}
	if app.Name != "" {
		db.Where("name = ?", app.Name)
	}

	if app.Status != 0 {
		db.Where("status = ?", app.Status)
	}
	db.Count(&total)
	if pageSize > 0 {
		appPage.PageSize = pageSize
		db.Limit(pageSize)
	} else {
		appPage.PageSize = 1000
		db.Limit(1000)
	}
	if page > 1 {
		appPage.Page = page
		db.Offset((page - 1) * pageSize)
	} else {
		appPage.Page = 1
		db.Offset(0)
	}

	if total != 0 {
		if err := DB.Find(&apps).Error; err != nil {
			return appPage, err
		}
	}

	appPage.Data = apps
	appPage.Total = int(total)
	return appPage, nil
}
