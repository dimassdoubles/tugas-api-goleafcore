package learngo

import (
	"fmt"

	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/controller/penjualan"
	"git.solusiteknologi.co.id/goleaf/apptemplate/pkg/learngo/controller/product"
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

	// product
	groupProduct := app.Group(config.ApiPrefix+"/v1/learngo/product", config.Middleware)
	groupProduct.Get("/get-list/:limit/:offset", product.GetProductList)
	groupProduct.Post("/add", product.AddProduct)
	groupProduct.Post("/edit", product.EditProduct)

	// penjualan
	groupPenjualan := app.Group(config.ApiPrefix+"/v1/learngo/penjualan", config.Middleware)
	groupPenjualan.Get("/get-list/:limit/:offset", penjualan.GetPenjualanList)

	return nil
}

func buildDefaultConfig(config Config) Config {
	return Config{
		ApiPrefix:  glutil.BlankOrDefault(config.ApiPrefix, glapi.API_PREFIX_DEFAULT),
		Middleware: glapi.OrEmptyMiddleware(config.Middleware),
	}
}
