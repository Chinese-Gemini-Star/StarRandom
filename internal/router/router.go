package router

import (
	"StarRandom/internal/controller"

	"github.com/kataras/iris/v12"
)

func BindRouters(app *iris.Application) {
	app.HandleDir("/", "./webapp")

	// ping
	app.Any("/ping", controller.PingController)

	// 生成随机数api
	app.Post("/postRandom", controller.PostRandomController)
}
