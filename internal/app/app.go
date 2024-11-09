package app

import "gorm.io/gorm"

type App struct {
	Config *Config
	DB     *gorm.DB
}

var app *App

func GetApp(configFilePath string) (*App, error) {
	if app != nil {
		return app, nil
	}
	app = &App{}
	if config, err := GetConfig(configFilePath); err != nil {
		return nil, err
	} else {
		app.Config = config
	}
	if db, err := LoadDB(app.Config); err != nil {
		return nil, err
	} else {
		app.DB = db
	}
	return app, nil
}
