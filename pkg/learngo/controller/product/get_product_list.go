package product

import (
	"context"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldata"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/gofiber/fiber/v2"
)

type OutGetProductList struct {
	ProductList []*tables.Product `json:"productList"`
}

type ParamGetProduct struct {
	Limit  int64 `json:"limit" example:"20" validate:"min=1"`
	Offset int64 `json:"page" example:"0" validate:"min=0"`
}

// GetProductList godoc
// @Router      /v1/learngo/product/get-list/:limit/:offset [get]
// @Summary     Mendapatkan list produk
// @Description Tidak ada
// @Tags        Learngo / Product
// @Accept      */*
// @Produce     json
// @Param       limit path int false "contoh: 10"
// @Param       offset path int false "contoh: 05"
// @Param       keyword query string false "contoh: hai"
// @Success     200      {object} OutGetProductList "Sample response"
// @Failure     500      {object} glapi.ResultFail "errorCode = unexpected.error"
func GetProductList(fc *fiber.Ctx) error {
	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		param := ParamGetProduct{}
		query := gldata.QueryStdGetList{}

		err := glapi.FetchValidAll(fc, &param, &query)
		if err != nil {
			return err
		}

		out := OutGetProductList{
			ProductList: make([]*tables.Product, 0),
		}

		err = gldb.SelectQMt(mt, *gldb.NewQBuilder().
			Add(" SELECT * ").
			Add(" FROM ", tables.PRODUCT, " ").
			Add(" WHERE true ").
			AddILike(" AND ", "keyword", query.Keyword, "product_code", "product_name").
			Add(" LIMIT :limit OFFSET :offset ").
			SetParam("limit", param.Limit).
			SetParam("offset", param.Offset),
			&out.ProductList)
		if err != nil {
			return err
		}

		return out
	})
}
