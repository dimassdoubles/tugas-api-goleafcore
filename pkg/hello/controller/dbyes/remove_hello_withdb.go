package dbyes

import (
	"context"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldata"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/gofiber/fiber/v2"
)

type BodyRemoveHello struct {
	Id int64 `json:"id" validate:"required" example:"10"`
}

type OutRemoveHello struct {
	Id   int64  `json:"id" example:"10"`
	Code string `json:"code" example:"H001"`
	Name string `json:"name" example:"Hellow"`
}

func RemoveHello(fc *fiber.Ctx) error {
	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		body := BodyRemoveHello{}
		err := glapi.FetchValidBody(fc, &body)
		if err != nil {
			return err
		}

		out := OutRemoveHello{}

		err = gldb.SelectRowQMt(mt, *gldb.NewQBuilder().
			Add(" DELETE FROM ", tables.S_HELLO).
			Add(" WHERE id = :id ").
			Add(" RETURNING id, code, name ").
			SetParam("id", body.Id), &out)
		if err != nil {
			return err
		}

		return out
	})
}
