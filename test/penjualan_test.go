package test_test

import (
	"log"
	"net/http"
	"testing"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/controller/penjualan"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

func TestPenjualan(t *testing.T) {
	gltest.TestApi(t, func(app *fiber.App, tx pgx.Tx) error {
		learngo.Setup(app, learngo.Config{})

		return nil

	}, func(assert *gltest.Assert, app *fiber.App, tx pgx.Tx, i int) interface{} {

		// GET LIST
		response := gltest.TestGet(t, app, "/api/v1/learngo/penjualan/get-list/10/0?dateFrom=20230101&dateTo="+glutil.DateNow())
		assert.AssertEquals(http.StatusOK, response.StatusCode, "Harusnya oke")

		var out goleafcore.Dto
		gltest.FetchRespBody(t, response, &out)
		log.Println("Output get list: ", out.PrettyString())

		// ADD
		itemList := []*penjualan.ItemPenjualan{
			{
				ProductId: 1,
				Qty:       decimal.NewFromInt(5),
				Price:     decimal.NewFromInt(3000),
			},
			{
				ProductId: 2,
				Qty:       decimal.NewFromInt(2),
				Price:     decimal.NewFromInt(3000),
			},
		}
		responseAdd := gltest.TestPost(t, app, "/api/v1/learngo/penjualan/add", penjualan.BodyAddPenjualan{
			TotalPembayaran: decimal.NewFromInt(50000),
			TotalKembalian:  decimal.NewFromInt(20000),
			ItemList:        itemList,
		})
		assert.AssertEquals(http.StatusOK, responseAdd.StatusCode, "Harusnya oke")

		gltest.FetchRespBody(t, responseAdd, &out)
		log.Println("Output add: ", out.PrettyString())

		return nil
	})
}
