package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	h "net/http"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/mbvisti/integration-test-in-go/config"
	"github.com/mbvisti/integration-test-in-go/repository/psql"
	"github.com/mbvisti/integration-test-in-go/server/http"
	"github.com/mbvisti/integration-test-in-go/service"
)

func TestIntegration_UserHandler_CreateUser(t *testing.T) {
	cfg := config.NewConfig()
	storage := psql.NewStorage()

	err := psql.RunUpMigrations(*cfg)
	if err != nil {
		t.Errorf("test setup failed for: CreateUser, with err: %v", err)
		return
	}

	userService := service.NewUser(storage)
	weightService := service.NewWeight(storage)

	userHandler := http.NewUserHandler(userService)
	weightHandler := http.NewWeightHandler(weightService)

	srv := http.NewHttp(cfg, *userHandler, *weightHandler)

	srv.SetupRoutes()
	r := srv.GetRouter()

	req := http.NewUserRequest{
		Name:          "Test user",
		Sex:           "male",
		WeightGoal:    "80",
		Email:         "test@gmail.com",
		Age:           99,
		Height:        185,
		ActivityLevel: 1,
	}
	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(req)
	if err != nil {
		log.Fatal(err)
	}
	rq, err := h.NewRequest(h.MethodPost, "/api/user", &buf)
	if err != nil {
		t.Error(err)
	}
	rq.Header.Add("Content-Type", "application/json")

	res, err := r.Test(rq, -1)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error(errors.New("create user endpoint did not return 200"))
	}

	// query the database to verify that a user was created based on the request
	// we sent
	newUser, err := storage.GetUserFromEmail(req.Email)
	if err != nil {
		t.Error(err)
	}

	if newUser.Height != req.Height {
		t.Error(errors.New("create user endpoint did not create user with correct details"))
	}
	if newUser.Name != req.Name {
		t.Error(errors.New("create user endpoint did not create user with correct details"))
	}
	if newUser.Sex != req.Sex {
		t.Error(errors.New("create user endpoint did not create user with correct details"))
	}
	if newUser.WeightGoal != req.WeightGoal {
		t.Error(errors.New("create user endpoint did not create user with correct details"))
	}
	if newUser.Email != req.Email {
		t.Error(errors.New("create user endpoint did not create user with correct details"))
	}
	if newUser.Age != req.Age {
		t.Error(errors.New("create user endpoint did not create user with correct details"))
	}
	if newUser.ActivityLevel != req.ActivityLevel {
		t.Error(errors.New("create user endpoint did not create user with correct details"))
	}

	t.Cleanup(func() {
		err := psql.RunDownMigrations(*cfg)
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				return
			}
			t.Errorf("test cleanup failed for: CreateUser endpoint, with err: %v", err)
		}
	})
}
