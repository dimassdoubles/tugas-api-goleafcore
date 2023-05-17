package product

import (
	"context"
	"errors"

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

// RemoveProduct godoc
// @Router      /v1/learngo/product/remove [delete]
// @Summary     Menghapus produk berdasarkan id
// @Description Id produk harus terdaftar
// @Tags        Learngo / Product
// @Accept      */*
// @Produce     json
// @Param       body body BodyRemoveProduct true "Sample JSON format"
// @Success     200      {object} OutRemoveProduct "Sample response"
// @Failure     500      {object} glapi.ResultFail500 "Unexpected Error. errorCode = unexpected.error, errorArgs: [message1,message2,...]"
func RemoveProduct(fc *fiber.Ctx) error {
	// validasi
	// - productId harus valid

	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		body := BodyRemoveProduct{}
		err := glapi.FetchValidBody(fc, &body)
		if err != nil {
			return err
		}

		out := OutRemoveProduct{}

		// productId harus valid
		var products []*tables.Product
		err = gldb.SelectQMt(mt, *gldb.NewQBuilder().
			Add(" SELECT * FROM ", tables.PRODUCT, " ").
			Add(" WHERE product_id = :productId").
			SetParam("productId", body.ProductId),
			&products,
		)
		if err != nil {
			return err
		}

		if len(products) == 0 {
			return errors.New("Produk id tidak valid")
		}

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
