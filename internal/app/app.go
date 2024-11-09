package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Config *Config
	DB     *gorm.DB
	Web    *gin.Engine
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
	app.Web = gin.Default()
	return app, nil
}

func (app *App) LinkStart() error {
	return app.Web.Run(app.Config.WEB.Port)
}
