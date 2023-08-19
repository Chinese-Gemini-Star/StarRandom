package main

import (
	server "StarRandom/server/controller"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	server.BindRouters(app)

	err := app.Listen(":24455")
	if err != nil {
		app.Logger().Error(err)
	}
}
