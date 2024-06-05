package api

import (
	"flag"
	"spendaro-api/pkg/util"
)

func BootstrapAndRunServer() {

	env := flag.String("APP_ENV", "development", "application runtime environment")
	debug := flag.Bool("DEBUG", false, "sets log level to debug if true")
	flag.Parse()

	util.LoadEnv()
	util.ReadAppConfig(env)
	util.SetupLogger(debug)

	RunServer()
}
