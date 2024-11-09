package main

import (
	"flag"
	"log/slog"
	a "myAccount/internal/app"
)

func main() {
	configFilePath := flag.String("config", "config.yaml", "config file path")
	flag.Parse()
	if app, err := a.GetApp(*configFilePath); err != nil {
		slog.Error("create app error:" + err.Error())
	} else {
		slog.Info("create app success:" + app.Config.APP.Name)
	}
}
