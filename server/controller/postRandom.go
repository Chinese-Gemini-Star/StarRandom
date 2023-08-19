package server

import (
	server "StarRandom/server/model"
	"github.com/kataras/iris/v12"
	"math/rand"
)

func postRandomController(ctx iris.Context) {
	// {"left":"1","right":"10","number":"6","isRepeat":"true"}
	randomInfo := &server.RandomInfo{}

	// 获取随机数要求
	err := ctx.ReadJSON(randomInfo)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
	ctx.Application().Logger().Printf("随机数范围:[%d:%d], 数量: %d, 是否可重:%t", randomInfo.Left, randomInfo.Right, randomInfo.Number, randomInfo.IsRepeat)

	randoms := server.Randoms{}

	// 生成随机数
	if randomInfo.IsRepeat {
		// 允许重复,直接生成即可
		for i := 0; i < randomInfo.Number; i++ {
			randoms.Randoms = append(randoms.Randoms, rand.Intn(randomInfo.Right-randomInfo.Left+1)+randomInfo.Left)
		}
	} else {
		// 禁止重复,借助map去重
		set := make(map[int]bool)
		for i := 0; i < randomInfo.Number; i++ {
			random := rand.Intn(randomInfo.Right-randomInfo.Left+1) + randomInfo.Left
			if _, exists := set[random]; exists {
				// 重复,重新生成
				i--
			} else {
				// 未重复
				randoms.Randoms = append(randoms.Randoms, random)
				set[random] = true
			}
		}
	}

	err = ctx.JSON(randoms)
	if err != nil {
		ctx.Application().Logger().Error(err)
	}
}
