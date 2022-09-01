package controller

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"github.com/gofiber/fiber/v2"
)

type ParamGetHello struct {
	Word string `json:"word" validate:"required"`
}

type OutGetHello struct {
	Word string `json:"word"`
}

func HelloGet(fc *fiber.Ctx) error {
	var param ParamGetHello
	err := glapi.FetchValidParams(fc, &param)
	if err != nil {
		return err
	}

	out := OutGetHello{
		Word: param.Word,
	}

	return glapi.Ok(fc, out)
}
