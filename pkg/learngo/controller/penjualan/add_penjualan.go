package penjualan

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldata"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gldb"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

type BodyAddPenjualan struct {
	TotalPembayaran decimal.Decimal `json:"totalPembayaran"  example:"20000"`
	TotalKembalian  decimal.Decimal `json:"totalKembalian" example:"30000"`
	ItemList        []ItemPenjualan `json:"itemList"`
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

// AddPenjualan godoc
// @Router      /v1/learngo/penjualan/add [post]
// @Summary     Menambahkan data penjualan baru
// @Description Tidak ada
// @Tags        Learngo / Penjualan
// @Accept      json
// @Produce     json
// @Param       body body BodyAddPenjualan true "Sample JSON format"
// @Success     200      {object} OutAddPenjualan "OK"
// @Failure     500      {object} glapi.ResultFail "errorCode = unexpected.error"
func AddPenjualan(fc *fiber.Ctx) error {
	// validasi
	// - productId harus valid
	// - qty harus lebih dari 0

	return glapi.ApiStd(fc, func(mt context.Context, audit *gldata.AuditData) interface{} {
		body := BodyAddPenjualan{}
		err := glapi.FetchValidBody(fc, &body)
		if err != nil {
			return err
		}

		out := OutAddPenjualan{}

		productIds := ""

		totalPenjualan := decimal.NewFromInt(0)

		err = gldb.BeginTrxMt(mt, func(tx pgx.Tx) error {

			for _, penjualanItem := range body.ItemList {
				// qty harus lebih dari 0
				if penjualanItem.Qty.LessThanOrEqual(decimal.NewFromInt(0)) {
					return errors.New("Minimal qty adalah 1")
				}

				// productId harus valid
				var products []*tables.Product
				err = gldb.SelectQTx(tx, *gldb.NewQBuilder().
					Add(" SELECT * FROM ", tables.PRODUCT, " ").
					Add(" WHERE product_id = :productId").
					SetParam("productId", penjualanItem.ProductId),
					&products,
				)
				if err != nil {
					return err
				}

				if len(products) == 0 {
					return errors.New(fmt.Sprintf("Produk id %v tidak valid", penjualanItem.ProductId))
				}

				totalPenjualan = totalPenjualan.Add(penjualanItem.Price.Mul(penjualanItem.Qty))
				productIds = productIds + fmt.Sprintf("%v, ", penjualanItem.ProductId)
			}

			err = gldb.SelectRowQTx(tx, *gldb.NewQBuilder().
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

			err = gldb.ExecQTx(tx, *gldb.NewQBuilder().
				Add(" INSERT INTO ", tables.PENJUALAN_ITEM, " ").
				Add(" (penjualan_id, product_id, qty, price, version) ").
				Add(" VALUES ").
				Add(" ", valuesInsert, " ").
				Log("Query insert penjualan item: "),
			)

			if err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return err
		}

		return out
	})
}
