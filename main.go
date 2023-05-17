package main

import (
	"git.solusiteknologi.co.id/goleaf/apptemplate/app/config"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/sirupsen/logrus"
)

// @title LEARN GO API
// @version 1.0
// @description APi documentation learn go
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @host https://myapp.id
// @BasePath /api/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @tokenUrl https://api.myapp.id/api/v1/auth/login
func main() {
	app := glinit.InitAll(glinit.LogConfig{
		LogFile: glutil.GetEnv(glinit.ENV_LOG_FILE, "./log/apptemplate.log"),
	}, glinit.DbConfig{
		Host:            glutil.GetEnv(glinit.ENV_DB_HOST, "172.17.0.1"),
		Port:            glutil.GetEnvInt(glinit.ENV_DB_PORT, 5432),
		Name:            glutil.GetEnv(glinit.ENV_DB_NAME, "myappdb"),
		User:            glutil.GetEnv(glinit.ENV_DB_USER, "sts"),
		Password:        glutil.GetEnv(glinit.ENV_DB_PASSWORD, "Awesome123!"),
		ApplicationName: "AppTemplateBackend",
	}, glinit.ServerConfig{
		Port: glutil.GetEnvInt(glinit.ENV_SERVER_PORT, glinit.DEFAULT_PORT),
	})

	config.ConfigureFiber(app)

	logrus.Info("Starting server")

	glinit.StartServer()

}
