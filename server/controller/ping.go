package server

import "github.com/kataras/iris/v12"

func pingController(ctx iris.Context) {
	_, err := ctx.WriteString("pong")
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}
