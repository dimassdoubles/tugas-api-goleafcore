package main

import (
	"git.solusiteknologi.co.id/goleaf/apptemplate/app/config"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glinit"
	"git.solusiteknologi.co.id/goleaf/goleafcore/glutil"
	"github.com/sirupsen/logrus"
)

func main() {
	app := glinit.InitAll(glinit.LogConfig{
		LogFile: glutil.GetEnv(glinit.ENV_LOG_FILE, "./log/apptemplate.log"),
	}, glinit.DbConfig{
		Host:            glutil.GetEnv(glinit.ENV_DB_HOST, "172.17.0.1"),
		Port:            glutil.GetEnvInt(glinit.ENV_DB_PORT, 5432),
		Name:            glutil.GetEnv(glinit.ENV_DB_NAME, "erp"),
		User:            glutil.GetEnv(glinit.ENV_DB_USER, "sts"),
		Password:        glutil.GetEnv(glinit.ENV_DB_PASSWORD, "Awesome123!"),
		ApplicationName: "AppTemplateBackend",
	}, glinit.ServerConfig{
		Port: glutil.GetEnvInt(glinit.ENV_SERVER_PORT, glinit.DEFAULT_PORT),
	})

	// glinit.InitLog()
	// app := glinit.InitServer(glinit.ServerConfig{
	// 	Port: 8000,
	// })

	config.ConfigureFiber(app)

	logrus.Info("Starting server")

	glinit.StartServer()

}
