package server

import (
	server "StarRandom/server/model"
	"github.com/kataras/iris/v12"
	"math/rand"
)

func postRandomController(ctx iris.Context) {
	// {"left":"1","right":"10","number":"6","isRepeat":"true"}

	randomInfo := &server.RandomInfo{}

	err := ctx.ReadJSON(randomInfo)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
	ctx.Application().Logger().Printf("随机数范围:[%d:%d], 数量: %d", randomInfo.Left, randomInfo.Right, randomInfo.Number)

	randoms := server.Randoms{}
	if randomInfo.IsRepeat {
		for i := 0; i < randomInfo.Number; i++ {
			randoms.Randoms = append(randoms.Randoms, rand.Intn(randomInfo.Right-randomInfo.Left+1)+randomInfo.Left)
		}
	} else {
		set := make(map[int]bool)
		for i := 0; i < randomInfo.Number; i++ {
			random := rand.Intn(randomInfo.Right-randomInfo.Left+1) + randomInfo.Left
			if _, exists := set[random]; !exists {
				randoms.Randoms = append(randoms.Randoms, random)
			} else {
				i--
			}
		}
	}

	err = ctx.JSON(randoms)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}
