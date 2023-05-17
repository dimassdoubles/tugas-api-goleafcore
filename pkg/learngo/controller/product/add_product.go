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

type BodyAddProduct struct {
	ProductCode string          `json:"productCode" validate:"required" example:"P001"`
	ProductName string          `json:"productName" validate:"required" example:"Indomie Rendang"`
	Price       decimal.Decimal `json:"price" validate:"required" example:"5000"`
}

type OutAddProduct struct {
	ProductId   int64           `json:"productId" example:"10"`
	ProductCode string          `json:"productCode" example:"P001"`
	ProductName string          `json:"productName" example:"Indomie Rendang"`
	Price       decimal.Decimal `json:"price" example:"5000"`
	Version     int64           `json:"version" exapmle:"0"`
}

// AddProduct godoc
// @Router      /v1/learngo/product/add [post]
// @Summary     Menambahkan data produk baru
// @Description Tidak ada
// @Tags        Learngo / Product
// @Accept      json
// @Produce     json
// @Param       body body BodyAddProduct true "Sample JSON format"
// @Success     200      {object} OutAddProduct "OK"
// @Failure     500      {object} glapi.ResultFail "errorCode = unexpected.error"
func AddProduct(fc *fiber.Ctx) error {
	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		body := BodyAddProduct{}
		err := glapi.FetchValidBody(fc, &body)
		if err != nil {
			return err
		}

		if body.Price.LessThanOrEqual(decimal.NewFromInt(0)) {
			return errors.New("'price' is required")
		}

		out := OutAddProduct{}

		err = gldb.SelectRowQMt(mt, *gldb.NewQBuilder().
			Add(" INSERT INTO ", tables.PRODUCT, " ").
			Add(" (product_code, product_name, price, version) ").
			Add(" VALUES ").
			Add(" (:productCode, :productName, :price, :version) ").
			Add(" RETURNING product_id, product_code, product_name, price, version ").
			SetParam("productCode", body.ProductCode).
			SetParam("productName", body.ProductName).
			SetParam("price", body.Price).
			SetParam("version", 0),
			&out)
		if err != nil {
			return err
		}

		return out
	})
}
