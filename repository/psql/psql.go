package psql

import (
	"errors"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mbvisti/integration-test-in-go/config"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage() *Storage {
	// side note: this would most likely go into a configuration package
	// that would expose a method to get these values and handle the cases
	// when they are not present. But for the sake of brevity, I leave them out.
	cfg := config.NewConfig()

	db, err := sqlx.Connect("postgres", cfg.GetDatabaseConnString())
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
