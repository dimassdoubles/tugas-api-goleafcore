package product

import (
	"context"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldata"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

type BodyRemoveProduct struct {
	ProductId int64 `json:"productId" validate:"required" example:"10"`
}

type OutRemoveProduct struct {
	ProductId   int64           `json:"productId" example:"10"`
	ProductCode string          `json:"productCode" example:"P001"`
	ProductName string          `json:"productName" example:"Indomie Rendang"`
	Price       decimal.Decimal `json:"price" example:"5000"`
	Version     int64           `json:"version" exapmle:"0"`
}

func RemoveProduct(fc *fiber.Ctx) error {
	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		body := BodyRemoveProduct{}
		err := glapi.FetchValidBody(fc, &body)
		if err != nil {
			return err
		}

		out := OutRemoveProduct{}

		err = gldb.SelectRowQMt(mt, *gldb.NewQBuilder().
			Add(" DELETE FROM ", tables.PRODUCT, " ").
			Add(" WHERE product_id = :productId ").
			Add(" RETURNING product_id, product_code, product_name, price, version ").
			SetParam("productId", body.ProductId).
			Log("Query remove product: "),
			&out)
		if err != nil {
			return err
		}

		return out
	})
}
