package hello

import (
	"fmt"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/controller/dbno"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glapi"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ApiPrefix  string
	Middleware fiber.Handler
}

func (c Config) String() string {
	return fmt.Sprint("groupRoute: ", c.ApiPrefix)
}

func Setup(app *fiber.App, config Config) error {
	config = buildDefaultConfig(config)
	logrus.Debug("Setup module hello with config : ", config)

	groupDbNo := app.Group(config.ApiPrefix+"/v1/hello/nodb", config.Middleware)
	groupDbNo.Get("/get-list/:count", dbno.GetHelloList)
	groupDbNo.Post("/add", dbno.AddHello)

	return nil
}

func buildDefaultConfig(config Config) Config {
	return Config{
		ApiPrefix:  glutil.BlankOrDefault(config.ApiPrefix, "/api"),
		Middleware: glapi.OrEmptyMiddleware(config.Middleware),
	}
}
