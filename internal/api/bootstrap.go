package api

import (
	"flag"
	"spendaro-api/pkg/util"
)

func BootstrapAndRunServer() {

	// Parse command line flags
	appEnv := flag.String("APP_ENV", "development", "application runtime environment")
	debug := flag.Bool("debug", false, "sets log level to debug if true")
	flag.Parse()

	util.LoadEnv()
	util.ReadAppConfig(appEnv)
	util.SetupLogger(debug)

	RunServer()
}
