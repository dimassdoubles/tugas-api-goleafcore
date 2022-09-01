package config

import (
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello"
	"github.com/gofiber/fiber/v2"
)

func ConfigureFiber(app *fiber.App) {

	// glauth.Setup(app, glauth.Config{})

	hello.Setup(app, hello.Config{
		ApiPrefix: "/api",
	})
}
