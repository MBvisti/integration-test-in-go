package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	databaseConfig *dbConfig
}

func NewConfig() *Config {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		panic(errors.New("env not set"))
	}
	dbConf, err := newDbConfig(env)
	if err != nil {
		panic(err)
	}
	return &Config{
		databaseConfig: dbConf,
	}
}

func (c *Config) GetDatabaseConnString() string {
	return c.GetDatabaseConnString()
}

type dbConfig struct {
	user     string
	password string
	host     string
	name     string
	driver   string
	port     int64
}

func newDbConfig(env string) (*dbConfig, error) {
	databasePort, err := strconv.ParseInt(
		os.Getenv("DB_PORT"), 0, 64,
	)
	if err != nil {
		return nil, errors.New("could not convert db port to int")
	}
	databaseHost := os.Getenv("DB_HOST")
	databaseUsername := os.Getenv("DB_USERNAME")
	databaseName := os.Getenv("DB_NAME")
	databasePassword := os.Getenv("DB_PASSWORD")
	databaseDriver := os.Getenv("DB_DRIVER")

	return &dbConfig{
		user:     databaseUsername,
		password: databasePassword,
		host:     databaseHost,
		name:     databaseName,
		driver:   databaseDriver,
		port:     databasePort,
	}, nil
}

func (d *dbConfig) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s"+
		" "+"dbname=%s  sslmode=disable",
		d.host, d.port, d.user,
		d.password, d.name,
	)
}
