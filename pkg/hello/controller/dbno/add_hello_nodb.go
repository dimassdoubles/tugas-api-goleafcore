package dbno

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"github.com/gofiber/fiber/v2"
)

type BodyGetHello struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}

func AddHello(fc *fiber.Ctx) error {
	var body BodyGetHello
	err := glapi.FetchValidBody(fc, &body)
	if err != nil {
		return err
	}

	return glapi.Ok(fc, body)
}
