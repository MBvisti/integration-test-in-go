package psql

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/mbvisti/integration-test-in-go/config"
)

type Storage struct {
	db *sql.DB
}

var (
	ErrNoNewMigrations = errors.New("no change")
)

func NewStorage() *Storage {
	// side note: this would most likely go into a configuration package
	// that would expose a method to get these values and handle the cases
	// when they are not present. But for the sake of brevity, I leave them out.
	cfg := config.NewConfig()

	db, err := sql.Open("postgres", cfg.GetDatabaseConnString())
	if err != nil {
		log.Printf("this is err: %+v", err)
		panic(errors.New("could not setup storage"))
	}
	log.Print("ping")
	err = db.Ping()
	if err != nil {
		panic(errors.New("could not ping db"))
	}
	log.Print("pong")
	return &Storage{
		db: db,
	}
}

func RunUpMigrations(cfg config.Config) error {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../migrations")
	migrationDir := filepath.Join("file://" + basePath)
	db, err := sql.Open("postgres", cfg.GetDatabaseConnString())
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.WithStack(err)
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(migrationDir, "postgres", driver)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, ErrNoNewMigrations) {
			return errors.WithStack(err)
		}
	}
	m.Close()
	return nil
}

func RunDownMigrations(cfg config.Config) error {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../migrations")
	migrationDir := filepath.Join("file://" + basePath)
	db, err := sql.Open("postgres", cfg.GetDatabaseConnString())
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.WithStack(err)
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(migrationDir, "postgres", driver)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := m.Down(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func DropEverythingInDatabase(cfg config.Config) error {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../migrations")
	migrationDir := filepath.Join("file://" + basePath)
	db, err := sql.Open("postgres", cfg.GetDatabaseConnString())
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.WithStack(err)
	}
	defer driver.Close()

	m, err := migrate.NewWithDatabaseInstance(migrationDir, "postgres", driver)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := m.Drop(); err != nil {
		return errors.WithStack(err)
	}
	srcErr, dbErr := m.Close()
	if srcErr != nil || dbErr != nil {
		return errors.Errorf("srcErr: %v and dbErr: %v", srcErr, dbErr)
	}

	return nil
}

func LoadFixtures(cfg config.Config) error {
	pathToFile := "/app/fixtures.sql"
	q, err := os.ReadFile(pathToFile)
	if err != nil {
		return errors.WithStack(err)
	}

	db, err := sql.Open("postgres", cfg.GetDatabaseConnString())
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = db.Exec(string(q))
	if err != nil {
		return errors.WithStack(err)
	}
	err = db.Close()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
