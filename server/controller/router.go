package server

import "github.com/kataras/iris/v12"

func BindRouters(app *iris.Application) {
	app.HandleDir("/", "./webapp")

	app.Any("/ping", func(ctx iris.Context) {
		_, err := ctx.WriteString("pong")
		if err != nil {
			app.Logger().Error(err)
		}
	})

	app.Post("postRandom", nil)
}
