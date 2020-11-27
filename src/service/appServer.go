package service

import "github.com/doter1995/user_center/src/model"

func (a App) FindApps(pageSize int, page int) (model.AppPagination, error) {
	app := model.App{
		Name:   a.Name,
		ID:     a.ID,
		Status: a.Status,
		Info:   a.Info,
		Path:   a.Path,
	}
	return model.FindApps(pageSize, page, app)
}

func (a App) CreateApp() (int, error) {
	var app = model.App{
		Name:   a.Name,
		Info:   a.Info,
		Path:   a.Path,
		Status: a.Status,
		Code:   "",
	}
	app1 := model.CreateApp(app)
	if app1.ID == 0 {
		return -1, nil
	}
	return 0, nil
}
