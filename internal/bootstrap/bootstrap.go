package bootstrap

import (
	"flag"
	"spendaro-api/internal/api"
	"spendaro-api/pkg/util"
)

func BootstrapServer() {

	appEnv, debug := readServerFlags()

	util.LoadEnv()
	util.ReadAppConfig(appEnv)
	util.SetupLogger(debug)
	api.NewEchoServer()

}

func readServerFlags() (*string, *bool) {
	appEnv := flag.String("APP_ENV", "dev", "application runtime environment")
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	return appEnv, debug
}
