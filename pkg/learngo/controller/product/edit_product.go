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

type BodyEditProduct struct {
	ProductId   int64           `json:"productId" validate:"required" example:"10"`
	ProductName string          `json:"productName" validate:"required" example:"Indomie Rendang"`
	Price       decimal.Decimal `json:"price" validate:"required" example:"5000"`
	Version     int64           `json:"version" example:"0"`
}

type OutEditProduct struct {
	ProductId   int64           `json:"productId" example:"10"`
	ProductCode string          `json:"productCode" example:"P001"`
	ProductName string          `json:"productName" example:"Indomie Rendang"`
	Price       decimal.Decimal `json:"price" example:"5000"`
	Version     int64           `json:"version" exapmle:"0"`
}

// EditProduct godoc
// @Router      /v1/learngo/product/edit [post]
// @Summary     Mengedit nama dan harga produk
// @Description Tidak ada
// @Tags        Learngo / Product
// @Accept      json
// @Produce     json
// @Param       body body BodyEditProduct true "Sample JSON format"
// @Success     200      {object} OutEditProduct "OK"
// @Failure     500      {object} glapi.ResultFail "errorCode = unexpected.error"
func EditProduct(fc *fiber.Ctx) error {
	// validasi
	// - productId harus valid
	// - version harus = version data sebelum diedit

	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		body := BodyEditProduct{}
		err := glapi.FetchValidBody(fc, &body)
		if err != nil {
			return err
		}

		out := OutEditProduct{}

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

		// version harus = version data sebelum diedit
		if body.Version != products[0].Version {
			return errors.New("Versi produk tidak sesuai dengan sistem")
		}

		err = gldb.SelectRowQMt(mt, *gldb.NewQBuilder().
			Add(" UPDATE ", tables.PRODUCT, " ").
			Add(" SET product_name = :productName, price = :price, version = :version+1 ").
			Add(" WHERE product_id = :productId  AND version = :version").
			Add(" RETURNING product_id, product_code, product_name, price, version ").
			SetParam("productId", body.ProductId).
			SetParam("productName", body.ProductName).
			SetParam("price", body.Price).
			SetParam("version", body.Version).
			Log("Query edit product: "),
			&out)
		if err != nil {
			return err
		}

		return out
	})
}
