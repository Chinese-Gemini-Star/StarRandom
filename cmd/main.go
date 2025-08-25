package main

import (
	"StarRandom/internal/router"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.Default()
	app.Logger().SetLevel("debug")

	router.BindRouters(app)

	err := app.Listen(":24455")
	if err != nil {
		app.Logger().Error(err)
	}
}
