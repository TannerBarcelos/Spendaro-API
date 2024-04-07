package bootstrap

import (
	"flag"
	"spendaro-api/internal/api"
	"spendaro-api/pkg/util"
)

func BootstrapServer() {

	appEnv, debug := readServerFlags()

    util.SetupLogger(debug)
	util.LoadEnv()
	util.ReadAppConfig(appEnv)
	api.NewServer()

}

func readServerFlags() (*string, *bool) {
	appEnv := flag.String("APP_ENV", "development", "application runtime environment")
	debug := flag.Bool("debug", false, "sets log level to debug if true")

	flag.Parse()

	return appEnv, debug
}
