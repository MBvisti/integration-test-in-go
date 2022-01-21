package psql_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
	"github.com/mbvisti/integration-test-in-go/config"
	"github.com/mbvisti/integration-test-in-go/entity"
	"github.com/mbvisti/integration-test-in-go/repository/psql"
)

// testing the happy path only - to improve upon these tests, we could consider
// using a table test
func TestIntegration_CreateUser(t *testing.T) {
	// create a NewStorage instance and run migrations
	cfg := config.NewConfig()
	storage := psql.NewStorage()

	err := psql.RunUpMigrations(*cfg)
	if err != nil {
		t.Errorf("test setup failed for: CreateUser, with err: %v", err)
		return
	}
	err = psql.LoadFixtures(*cfg)
	if err != nil {
		t.Errorf("test setup failed for: CreateUser, with err: %v", err)
		return
	}

	// run the test
	t.Run("should create a new user", func(t *testing.T) {
		newUser, err := entity.NewUser(
			"Jon Snow", "male", "90", "theyoungwolf@stark.com", 16, 182, 1)
		if err != nil {
			t.Errorf("failed to run CreateUser with error: %v", err)
			return
		}

		err = storage.CreateUser(*newUser)
		// assert there is no err
		if err != nil {
			t.Errorf("failed to create new user with err: %v", err)
			return
		}

		// now lets verify that the user is actually created using a
		// separate connection to the DB and pure sql
		db, err := sql.Open("postgres", cfg.GetDatabaseConnString())
		if err != nil {
			t.Errorf("failed to connect to database with err: %v", err)
			return
		}
		queryResult := entity.User{}
		err = db.QueryRow("SELECT id, name, email FROM users WHERE email=$1",
			"theyoungwolf@stark.com").Scan(
			&queryResult.ID, &queryResult.Name, &queryResult.Email,
		)
		if err != nil {
			t.Errorf("this was query err: %v", err)
			return
		}

		if queryResult.Name != newUser.Name {
			t.Error(`failed 'should create a new user' wanted name did not match 
				returned value`)
			return
		}
		if queryResult.Email != newUser.Email {
			t.Error(`failed 'should create a new user' wanted email did not match 
				returned value`)
			return
		}
	})

	// run some clean up, i.e. clean the database so we have a clean env
	// when we run the next test
	t.Cleanup(func() {
		err := psql.RunDownMigrations(*cfg)
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				return
			}
			t.Errorf("test cleanup failed for: CreateUser, with err: %v", err)
		}
	})
}
