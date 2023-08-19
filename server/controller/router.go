package server

import "github.com/kataras/iris/v12"

func BindRouters(app *iris.Application) {
	app.HandleDir("/", "./webapp")

	// ping
	app.Any("/ping", pingController)

	// 生成随机数api
	app.Post("postRandom", postRandomController)
}
