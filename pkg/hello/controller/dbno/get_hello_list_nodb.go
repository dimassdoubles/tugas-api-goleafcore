package dbno

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"github.com/gofiber/fiber/v2"
)

type ParamGetHello struct {
	Count int64 `json:"count" example:"5" validate:"min=2"`
}

type OutGetHelloList struct {
	HelloList []*Hello `json:"helloList"`
}

type Hello struct {
	Word string `json:"word" example:"Hello world"`
}

func GetHelloList(fc *fiber.Ctx) error {
	var param ParamGetHello
	err := glapi.FetchValidParams(fc, &param)
	if err != nil {
		return err
	}

	helloList := make([]*Hello, 0)

	for i := 0; i < int(param.Count); i++ {
		helloList = append(helloList, &Hello{
			Word: "Hello world without database",
		})
	}

	out := OutGetHelloList{
		HelloList: helloList,
	}

	return glapi.Ok(fc, out)
}
