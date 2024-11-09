package main

import (
	"flag"
	"log/slog"
	a "myAccount/internal/app"
	"myAccount/internal/router"
)

func main() {
	configFilePath := flag.String("config", "config.yaml", "config file path")
	flag.Parse()
	if app, err := a.GetApp(*configFilePath); err != nil {
		slog.Error("create app error:" + err.Error())
	} else {
		// 封装接口
		router.SetRouter(app)
		// 启动
		if err := app.LinkStart(); err != nil {
			slog.Error("app run error:" + err.Error())
		}
	}
}
