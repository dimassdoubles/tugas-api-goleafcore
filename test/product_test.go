package test_test

import (
	"log"
	"net/http"
	"testing"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/controller/product"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/shopspring/decimal"
)

func TestProductApi(t *testing.T) {
	gltest.TestApi(t, func(app *fiber.App, tx pgx.Tx) error {
		learngo.Setup(app, learngo.Config{})

		return nil

	}, func(assert *gltest.Assert, app *fiber.App, tx pgx.Tx, i int) interface{} {

		// GET LIST
		response := gltest.TestGet(t, app, "/api/v1/learngo/product/get-list/10/0")
		assert.AssertEquals(http.StatusOK, response.StatusCode, "Harusnya oke")

		var out goleafcore.Dto
		gltest.FetchRespBody(t, response, &out)
		log.Println("Output get list: ", out.PrettyString())

		// ADD
		responseAdd := gltest.TestPost(t, app, "/api/v1/learngo/product/add", product.BodyAddProduct{
			ProductCode: "P999",
			ProductName: "Indomie Goreng",
			Price:       decimal.NewFromInt(5000),
		})
		assert.AssertEquals(http.StatusOK, responseAdd.StatusCode, "Harusnya oke")

		gltest.FetchRespBody(t, responseAdd, &out)
		log.Println("Output add: ", out.PrettyString())

		return nil
	})
}
