package main

import (
	"log"

	"github.com/mbvisti/integration-test-in-go/repository/psql"
	"github.com/mbvisti/integration-test-in-go/service"
)

func main() {
	storage := psql.NewStorage()

	// these would then be passed to some handler in whatever way we choose
	// to expose these. It could be through REST, gRPC, graphQL...
	userService := service.NewUser(storage)
	weightService := service.NewWeight(storage)

	// services aren't really running - this is just done to shut up the
	// compiler. We don't really need them but leaving them here for
	// illustration
	log.Printf("service running: %v - %v", userService, weightService)
}
