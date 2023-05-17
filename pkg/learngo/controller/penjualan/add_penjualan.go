package penjualan

import (
	"context"
	"fmt"
	"log"
	"strings"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldata"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/gofiber/fiber/v2"
	"github.com/shopspring/decimal"
)

type BodyAddPenjualan struct {
	TotalPembayaran decimal.Decimal  `json:"totalPembayaran"  example:"20000"`
	TotalKembalian  decimal.Decimal  `json:"totalKembalian" example:"30000`
	ItemList        []*ItemPenjualan `json:"itemList" example:"[{\"productId\": 10, \"qty\": \"5\", \"price\": \"10000\"}, {\"productId\": 20, \"qty\": \"3\", \"price\": \"15000\"}]"`
}

type ItemPenjualan struct {
	ProductId int64           `json:"productId" example:"10"`
	Qty       decimal.Decimal `json:"qty" example:"5"`
	Price     decimal.Decimal `json:"price" example:"10000"`
}

type OutAddPenjualan struct {
	PenjualanId     int64           `json:"penjualanId" example:"10"`
	TotalPenjualan  decimal.Decimal `json:"totalPenjualan" example:"20000"`
	TotalPembayaran decimal.Decimal `json:"totalPembayaran" example:"50000"`
	TotalKembalian  decimal.Decimal `json:"totalKembalian" example:"30000"`
	Version         int64           `json:"version" exapmle:"0"`
}

func AddPenjualan(fc *fiber.Ctx) error {
	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		body := BodyAddPenjualan{}
		err := glapi.FetchValidBody(fc, &body)
		if err != nil {
			return err
		}

		out := OutAddPenjualan{}

		totalPenjualan := decimal.NewFromInt(0)

		for _, penjualanItem := range body.ItemList {
			totalPenjualan = totalPenjualan.Add(penjualanItem.Price.Mul(penjualanItem.Qty))
		}

		log.Println("Total Penjualan: ", totalPenjualan)

		err = gldb.SelectRowQMt(mt, *gldb.NewQBuilder().
			Add(" INSERT INTO ", tables.PENJUALAN, " ").
			Add(" (total_penjualan, total_pembayaran, total_kembalian, tanggal, version) ").
			Add(" VALUES ").
			Add(" (:totalPenjualan, :totalPembayaran, :totalKembalian, :tanggal, :version) ").
			Add(" RETURNING penjualan_id, total_penjualan, total_pembayaran, total_kembalian, version ").
			SetParam("totalPenjualan", totalPenjualan).
			SetParam("totalPembayaran", body.TotalPembayaran).
			SetParam("totalKembalian", body.TotalKembalian).
			SetParam("tanggal", glutil.DateNow()).
			SetParam("version", 0),
			&out)
		if err != nil {
			return err
		}

		valuesInsert := ""
		for _, penjualanItem := range body.ItemList {
			valuesInsert = valuesInsert + fmt.Sprintf("(%v, %v, %v, %v, 0),", out.PenjualanId, penjualanItem.ProductId, penjualanItem.Qty, penjualanItem.Price)
		}

		valuesInsert = strings.TrimRight(valuesInsert, ",") + ";"

		log.Println("values insert %s", valuesInsert)

		err = gldb.ExecQMt(mt, *gldb.NewQBuilder().
			Add(" INSERT INTO ", tables.PENJUALAN_ITEM, " ").
			Add(" (penjualan_id, product_id, qty, price, version) ").
			Add(" VALUES ").
			Add(" ", valuesInsert, " ").
			Log("Query insert penjualan item: "),
		)

		if err != nil {
			return err
		}

		return out
	})
}
