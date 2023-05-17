package config

import (
	_ "git.solusiteknologi.co.id/goleaf/apptemplate/docs"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo"
	"git.solusiteknologi.co.id/goleaf/glauth"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func ConfigureFiber(app *fiber.App) {

	// TODO jika mau menggunakan security api
	accessMd := glauth.MiddlewareTask
	glauth.Setup(app, glauth.Config{})

	hello.Setup(app, hello.Config{
		ApiPrefix: glapi.API_PREFIX_DEFAULT,
		// TODO jika mau menggunakan security api
		Middleware: accessMd,
	})

	learngo.Setup(app, learngo.Config{
		ApiPrefix: glapi.API_PREFIX_DEFAULT,
		// TODO jika mau menggunakan security api
		Middleware: accessMd,
	})

	// API docs
	app.Get("/api/docs/*", swagger.HandlerDefault)
	app.Get("/api/docs", func(c *fiber.Ctx) error {
		return c.Redirect("/api/docs/", fiber.StatusTemporaryRedirect)
	})
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("/api/docs/", fiber.StatusTemporaryRedirect)
	})
	app.Get("/docs/*", func(c *fiber.Ctx) error {
		return c.Redirect("/api/docs/", fiber.StatusTemporaryRedirect)
	})
}
