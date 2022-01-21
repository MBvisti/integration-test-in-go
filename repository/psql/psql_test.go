package psql_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/mbvisti/integration-test-in-go/config"
	"github.com/mbvisti/integration-test-in-go/repository/psql"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	TestStorage *psql.Storage
	TestDb      *sql.DB
	Cfg         *config.Config
}

func (s *RepositoryTestSuite) SetupSuite() {
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

	s.Cfg = cfg
	s.TestDb = db
	s.TestStorage = storage
}

func (s *RepositoryTestSuite) TearDownSuite() {
	err := psql.DropEverythingInDatabase(*s.Cfg)
	if err != nil {
		panic(err)
	}
}

func (s *RepositoryTestSuite) BeforeTest(suiteName, testName string) {
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

func (s *RepositoryTestSuite) AfterTest(suiteName, testName string) {
	err := psql.RunDownMigrations(*s.Cfg)
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func TestIntegration_RepositoryTestSuite(t *testing.T) {
	repoSuite := new(RepositoryTestSuite)
	suite.Run(t, repoSuite)
}
