package test_test

import (
	"log"
	"net/http"
	"testing"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func TestPenjualan(t *testing.T) {
	gltest.TestApi(t, func(app *fiber.App, tx pgx.Tx) error {
		learngo.Setup(app, learngo.Config{})

		return nil

	}, func(assert *gltest.Assert, app *fiber.App, tx pgx.Tx, i int) interface{} {

		// GET LIST
		response := gltest.TestGet(t, app, "/api/v1/learngo/penjualan/get-list/10/0?dateFrom=20230508&dateTo=20230511")
		assert.AssertEquals(http.StatusOK, response.StatusCode, "Harusnya oke")

		var out goleafcore.Dto
		gltest.FetchRespBody(t, response, &out)
		log.Println("Output get list: ", out.PrettyString())

		return nil
	})
}
