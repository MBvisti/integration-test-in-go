package http_test

import (
	"os"
	"testing"

	"github.com/mbvisti/integration-test-in-go/config"
	"github.com/mbvisti/integration-test-in-go/repository/psql"
)

// TestMain gets run before running any other _test.go files in each package
// here, we use it to make sure we start from a clean slate
func TestMain(m *testing.M) {
	cfg := config.NewConfig()
	// make sure we start from a clean slate
	err := psql.DropEverythingInDatabase(*cfg)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
