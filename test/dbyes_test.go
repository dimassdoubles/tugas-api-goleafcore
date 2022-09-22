package test_test

import (
	"log"
	"net/http"
	"testing"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/controller/dbyes"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/tables"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func TestDbYesApi(t *testing.T) {
	gltest.TestApi(t, func(app *fiber.App, tx pgx.Tx) error {
		hello.Setup(app, hello.Config{})

		return gltest.InitTableSample(tx, 20, tables.S_HELLO, "id", "code", "name")

	}, func(assert *gltest.Assert, app *fiber.App, tx pgx.Tx, i int) interface{} {

		// GET LIST
		response := gltest.TestGet(t, app, "/api/v1/hello/dbyes/get-list/10")
		assert.AssertEquals(http.StatusOK, response.StatusCode, "Harusnya oke")

		var out goleafcore.Dto
		gltest.FetchRespBody(t, response, &out)
		log.Println("Output get list: ", out.PrettyString())

		// ADD
		responseAdd := gltest.TestPost(t, app, "/api/v1/hello/dbyes/add", dbyes.BodyAddHello{
			Name: "Fredi",
			Code: "F001",
		})
		assert.AssertEquals(http.StatusOK, responseAdd.StatusCode, "Harusnya oke")

		gltest.FetchRespBody(t, responseAdd, &out)
		log.Println("Output add: ", out.PrettyString())

		// REMOVE
		responseRemove := gltest.TestDelete(t, app, "/api/v1/hello/dbyes/remove", dbyes.BodyRemoveHello{
			Id: out.GetInt64("id"),
		})
		assert.AssertEquals(http.StatusOK, responseRemove.StatusCode, "Harusnya oke")

		gltest.FetchRespBody(t, responseRemove, &out)
		log.Println("Output remove: ", out.PrettyString())

		return nil
	})
}
