package test_test

import (
	"log"
	"net/http"
	"testing"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/controller/dbno"
	"git.solusiteknologi.co.id/goleaf/goleafcore"
	"git.solusiteknologi.co.id/goleaf/goleafcore/gltest"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func TestDbnoApi(t *testing.T) {
	gltest.TestApi(t, func(app *fiber.App, tx pgx.Tx) error {
		hello.Setup(app, hello.Config{})

		return nil
	}, func(assert *gltest.Assert, app *fiber.App, tx pgx.Tx, i int) interface{} {

		response := gltest.TestGet(t, app, "/api/v1/hello/dbno/get-list/2")
		assert.AssertEquals(http.StatusOK, response.StatusCode, "Harusnya oke")

		var out goleafcore.Dto
		gltest.FetchRespBody(t, response, &out)
		log.Println("Output get list: ", out.PrettyString())

		responseAdd := gltest.TestPost(t, app, "/api/v1/hello/dbno/add", dbno.BodyGetHello{
			Name: "Fredi",
			Age:  100,
		})
		assert.AssertEquals(http.StatusOK, responseAdd.StatusCode, "Harusnya oke")

		gltest.FetchRespBody(t, responseAdd, &out)
		log.Println("Output add: ", out.PrettyString())

		return nil
	})
}
