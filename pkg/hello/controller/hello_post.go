package controller

import (
	"fmt"

	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"github.com/gofiber/fiber/v2"
)

type BodyGetHello struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}

func HelloPost(fc *fiber.Ctx) error {
	var body BodyGetHello
	err := glapi.FetchValidBody(fc, &body)
	if err != nil {
		return err
	}

	out := OutGetHello{
		Word: fmt.Sprint("Hello ", body.Name, " you are ", body.Age, " years"),
	}

	return glapi.Ok(fc, out)
}
