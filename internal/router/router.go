package router

import (
	"github.com/gin-gonic/gin"
	a "myAccount/internal/app"
	"myAccount/internal/utils/r"
)

func SetRouter(app *a.App) {
	web := app.Web
	public := web.Group("public")
	{
		// public
		hello(public)
	}
}

func hello(router *gin.RouterGroup) {
	router.GET("/hello", func(ctx *gin.Context) {
		r.JSON(ctx, r.OK("hello world"))
	})
}
