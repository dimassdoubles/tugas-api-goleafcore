package config

import (
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"github.com/gofiber/fiber/v2"
)

func ConfigureFiber(app *fiber.App) {

	// TODO jika mau menggunakan security api
	// accessMd := glauth.MiddlewareTask
	// glauth.Setup(app, glauth.Config{})

	hello.Setup(app, hello.Config{
		ApiPrefix: glapi.API_PREFIX_DEFAULT,
		// TODO jika mau menggunakan security api
		// Middleware: accessMd,
	})

	// call Setup other modules
	//
}
