package psql

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage() *Storage {
	// side note: this would most likely go into a configuration package
	// that would expose a method to get these values and handle the cases
	// when they are not present. But for the sake of brevity, I leave them out.
	databasePort, err := strconv.ParseInt(
		os.Getenv("DB_PORT"), 0, 64,
	)
	if err != nil {
		panic(errors.New("could not convert db port to int"))
	}
	databaseHost := os.Getenv("DB_HOST")
	databaseUsername := os.Getenv("DB_USERNAME")
	databaseName := os.Getenv("DB_NAME")
	databasePassword := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s"+
		" "+"dbname=%s sslmode=disable",
		databaseHost, databasePort, databaseUsername, databasePassword,
		databaseName,
	)

	db, err := sqlx.Connect("postgres", connectionString)
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
