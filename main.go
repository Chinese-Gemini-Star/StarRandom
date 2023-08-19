package main

import (
	server "StarRandom/server/controller"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()

	server.BindRouters(app)

	err := app.Listen(":80")
	if err != nil {
		app.Logger().Error(err)
	}
}
