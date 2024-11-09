package app

type App struct {
	Config *Config
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
	return app, nil
}
