package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	srvConfig      *serverConfig
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
	srvConf, err := newserverConfig()
	if err != nil {
		panic(err)
	}
	return &Config{
		databaseConfig: dbConf,
		srvConfig:      srvConf,
	}
}

func (c *Config) GetDatabaseConnString() string {
	return c.databaseConfig.getConnectionString()
}

func (c *Config) GetServerPort() int64 {
	return c.srvConfig.port
}
func (c *Config) GetServerReadTimeOut() time.Duration {
	return c.srvConfig.readTimeout
}
func (c *Config) GetServerWriteTimeOut() time.Duration {
	return c.srvConfig.writeTimeout
}

type dbConfig struct {
	user     string
	password string
	host     string
	name     string
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
	if databaseHost == "" {
		return nil, errors.New("databaseHost was empty")
	}
	databaseUsername := os.Getenv("DB_USERNAME")
	if databaseUsername == "" {
		return nil, errors.New("databaseUsername was empty")
	}
	databaseName := os.Getenv("DB_NAME")
	if databaseName == "" {
		return nil, errors.New("databaseName was empty")
	}
	databasePassword := os.Getenv("DB_PASSWORD")
	if databasePassword == "" {
		return nil, errors.New("databasePassword was empty")
	}
	return &dbConfig{
		user:     databaseUsername,
		password: databasePassword,
		host:     databaseHost,
		name:     databaseName,
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

type serverConfig struct {
	readTimeout  time.Duration
	writeTimeout time.Duration
	port         int64
}

func newserverConfig() (*serverConfig, error) {
	srvPort := os.Getenv("SERVER_PORT")
	if srvPort == "" {
		return nil, errors.New("missing server port env variable")
	}
	port, err := strconv.Atoi(srvPort)
	if err != nil {
		return nil, err
	}
	return &serverConfig{
		port:         int64(port),
		readTimeout:  time.Minute * 12,
		writeTimeout: time.Minute * 12,
	}, nil
}
