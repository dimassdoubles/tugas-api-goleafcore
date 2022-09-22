package hello

import (
	"fmt"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/controller/dbno"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/hello/controller/dbyes"
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
	return fmt.Sprint("apiPrefix: ", c.ApiPrefix)
}

func Setup(app *fiber.App, config Config) error {
	config = buildDefaultConfig(config)
	logrus.Debug("Setup module hello with config --> ", config)

	groupDbNo := app.Group(config.ApiPrefix+"/v1/hello/dbno", config.Middleware)
	groupDbNo.Get("/get-list/:count", dbno.GetHelloList)
	groupDbNo.Post("/add", dbno.AddHello)

	groupDbYes := app.Group(config.ApiPrefix+"/v1/hello/dbyes", config.Middleware)
	groupDbYes.Get("/get-list/:limit/:offset?", dbyes.GetHelloList)
	groupDbYes.Post("/add", dbyes.AddHello)
	groupDbYes.Delete("/remove", dbyes.RemoveHello)

	return nil
}

func buildDefaultConfig(config Config) Config {
	return Config{
		ApiPrefix:  glutil.BlankOrDefault(config.ApiPrefix, glapi.API_PREFIX_DEFAULT),
		Middleware: glapi.OrEmptyMiddleware(config.Middleware),
	}
}
