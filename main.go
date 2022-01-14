package main

import (
	"github.com/mbvisti/integration-test-in-go/config"
	"github.com/mbvisti/integration-test-in-go/repository/psql"
	"github.com/mbvisti/integration-test-in-go/server/http"
	"github.com/mbvisti/integration-test-in-go/service"
)

func main() {
	cfg := config.NewConfig()
	storage := psql.NewStorage()

	// these would then be passed to some handler in whatever way we choose
	// to expose these. It could be through REST, gRPC, graphQL...
	userService := service.NewUser(storage)
	weightService := service.NewWeight(storage)

	userHandler := http.NewUserHandler(userService)
	weightHandler := http.NewWeightHandler(weightService)

	srv := http.NewHttp(cfg, *userHandler, *weightHandler)
	srv.Start()
}
