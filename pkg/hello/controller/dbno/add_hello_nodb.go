package dbno

import (
	"fmt"

	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BodyGetHello struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"required"`
}

type OutGetHello struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Word string `json:"word"`
}

func AddHello(fc *fiber.Ctx) error {
	body := BodyGetHello{}
	err := glapi.FetchValidBody(fc, &body)
	if err != nil {
		return err
	}

	logrus.Debug("Body masuk? ", body)

	out := OutGetHello{
		Name: body.Name,
		Age:  body.Age,
		Word: fmt.Sprint("You are awesome [", body.Name, "]. Age ", body.Age, " is just a number"),
	}

	return glapi.Ok(fc, out)
}
