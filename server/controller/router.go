package server

import "github.com/kataras/iris/v12"

func BindRouters(app *iris.Application) {
	app.HandleDir("/", "./webapp")

	app.Any("/ping", pingController)

	app.Post("postRandom", postRandomController)
}
