package useauth

import (
	"context"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldata"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/gofiber/fiber/v2"
)

type OutGetHelloList struct {
	HelloList []*Hello `json:"helloList"`
}

type Hello struct {
	Id   int64  `json:"id" example:"10"`
	Code string `json:"code" example:"H001"`
	Name string `json:"name" example:"Hellow"`
}

func GetHelloList(fc *fiber.Ctx) error {
	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		param := gldata.ParamStdGetList{}
		query := gldata.QueryStdGetList{}

		err := glapi.FetchValidAll(fc, &param, &query)
		if err != nil {
			return err
		}

		out := OutGetHelloList{
			HelloList: make([]*Hello, 0),
		}

		err = gldb.SelectQMt(mt, *gldb.NewQBuilder().
			Add(" SELECT A.id, A.code, A.name  ").
			Add(" FROM ", tables.S_HELLO, " A ").
			Add(" WHERE true ").
			AddILike(" AND ", "keyword", query.Keyword, "A.code", "A.name").
			Add(" LIMIT :limit OFFSET :offset ").
			SetParam("limit", param.Limit).
			SetParam("offset", param.Page*param.Limit),
			&out.HelloList)
		if err != nil {
			return err
		}

		return out
	})
}
