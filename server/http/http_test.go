package http_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-migrate/migrate/v4"
	"github.com/mbvisti/integration-test-in-go/config"
	"github.com/mbvisti/integration-test-in-go/repository/psql"
	"github.com/mbvisti/integration-test-in-go/server/http"
	"github.com/mbvisti/integration-test-in-go/service"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type HttpTestSuite struct {
	suite.Suite
	TestStorage *psql.Storage
	TestDb      *sql.DB
	TestRouter  *fiber.App
	Cfg         *config.Config
}

func (s *HttpTestSuite) SetupSuite() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cfg := config.NewConfig()

	db, err := sql.Open("postgres", cfg.GetDatabaseConnString())
	if err != nil {
		panic(errors.WithStack(err))
	}

	err = db.Ping()
	if err != nil {
		panic(errors.WithStack(err))
	}
	storage := psql.NewStorage()

	userService := service.NewUser(storage)
	weightService := service.NewWeight(storage)

	userHandler := http.NewUserHandler(userService)
	weightHandler := http.NewWeightHandler(weightService)

	srv := http.NewHttp(cfg, *userHandler, *weightHandler)

	srv.SetupRoutes()
	r := srv.GetRouter()

	s.Cfg = cfg
	s.TestDb = db
	s.TestStorage = storage
	s.TestRouter = r
}

func (s *HttpTestSuite) TearDownSuite() {
	err := psql.DropEverythingInDatabase(*s.Cfg)
	if err != nil {
		panic(err)
	}
}

func (s *HttpTestSuite) BeforeTest(suiteName, testName string) {
	err := psql.RunUpMigrations(*s.Cfg)
	if err != nil && errors.Is(err, migrate.ErrNoChange) {
		log.Printf("this is the error: %+v", err)
		panic(errors.WithStack(err))
	}

	err = psql.LoadFixtures(*s.Cfg)
	if err != nil {
		log.Printf("this is the error: %+v", err)
		panic(errors.WithStack(err))
	}
}

func (s *HttpTestSuite) AfterTest(suiteName, testName string) {
	err := psql.RunDownMigrations(*s.Cfg)
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func TestIntegration_HttpTestSuite(t *testing.T) {
	repoSuite := new(HttpTestSuite)
	suite.Run(t, repoSuite)
}
