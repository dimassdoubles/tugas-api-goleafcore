package dbno

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"github.com/gofiber/fiber/v2"
)

type ParamGetHello struct {
	Count int64 `json:"count" example:"5" validate:"required min=2"`
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

	out := OutGetHelloList{
		HelloList: make([]*Hello, param.Count),
	}
	for _, hello := range out.HelloList {
		hello.Word = "Hello world no db"
	}

	return glapi.Ok(fc, out)
}
