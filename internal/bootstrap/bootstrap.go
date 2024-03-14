package bootstrap

import (
	"log"
	"spendaro-api/internal/api/server"
	"spendaro-api/pkg/util"
)

// BootstrapServer bootstraps the server and starts it
func BootstrapServer() {
	log.Println("Bootstrapping server") // TODO: Implement slog for structured logging (so we can pass logs to Grafana Loki for example)
	util.LoadEnv()
	server.NewEchoServer(":1323")
}
