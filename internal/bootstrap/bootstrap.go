package bootstrap

import (
	"log"
	"spendaro-api/internal/api"
	"spendaro-api/pkg/util"
)

func BootstrapServer() {
	util.LoadEnv()
	util.ReadAppConfig()
	log.Println("Bootstrapping server") // TODO: Implement slog for structured logging (so we can pass logs to Grafana Loki for example)
	api.NewEchoServer()
}
