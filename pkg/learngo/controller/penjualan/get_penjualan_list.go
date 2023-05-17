package penjualan

import (
	"context"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldata"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

type OutGetPenjualanList struct {
	PenjualanList []*Penjualan `json:"productList"`
}

type Penjualan struct {
	PenjualanId     int64           `json:"penjualanId" example:"10"`
	TotalPenjualan  decimal.Decimal `json:"totalPenjualan" example:"120000"`
	TotalPembayaran decimal.Decimal `json:"totalPembayaran" example:"150000"`
	TotalKembalian  decimal.Decimal `json:"totalKembalian" example:"30000"`
	Version         int64           `json:"version" example:"0"`
	ItemCount       int64           `json:"itemCount" example:"13"`
}

type QueryGetPenjualanList struct {
	DateFrom string `json:"dateFrom" example:"20230115"`
	DateTo   string `json:"dateTo" example:"20230120"`
}

type ParamGetPenjualanList struct {
	Limit  int64 `json:"limit" example:"20" validate:"min=1"`
	Offset int64 `json:"page" example:"0" validate:"min=0"`
}

// GetPenjualanList godoc
// @Router      /api/v1/learngo/penjualan/get-list/:limit/:offset [get]
// @Summary     Mendapatkan daftar penjualan
// @Description Tidak ada
// @Tags        Learngo / Penjualan
// @Accept      */*
// @Produce     json
// @Param       limit path int false "contoh: 10"
// @Param       offset path int false "contoh: 05"
// @Param       query query QueryGetPenjualanList false "Sample JSON format"
// @Success     200      {object} OutGetPenjualanList "Sample response"
// @Failure     500      {object} glapi.ResultFail "errorCode = unexpected.error"
func GetPenjualanList(fc *fiber.Ctx) error {
	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		param := ParamGetPenjualanList{}
		query := QueryGetPenjualanList{}

		err := glapi.FetchValidAll(fc, &param, &query)
		if err != nil {
			return err
		}

		out := OutGetPenjualanList{
			PenjualanList: make([]*Penjualan, 0),
		}

		err = gldb.SelectQMt(mt, *gldb.NewQBuilder().
			Add(" SELECT ").
			Add(" A.penjualan_id, A.total_penjualan, A.total_pembayaran, A.total_kembalian, version, ").
			Add(" (SELECT COUNT(1) FROM ", tables.PENJUALAN_ITEM, " B WHERE B.penjualan_id = A.penjualan_id) AS item_count ").
			Add(" FROM ").
			Add(tables.PENJUALAN, " A ").
			Add(" where ").
			Add(" A.tanggal >= :dateFrom AND").
			Add(" A.tanggal  <= :dateTo ").
			Add(" LIMIT :limit OFFSET :offset ").
			SetParam("dateFrom", query.DateFrom).
			SetParam("dateTo", query.DateTo).
			SetParam("limit", param.Limit).
			SetParam("offset", param.Offset).Log("query get penjualan list"),

			&out.PenjualanList)
		if err != nil {
			return err
		}

		return out
	})
}
