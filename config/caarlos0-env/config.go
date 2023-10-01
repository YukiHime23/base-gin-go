package config_caarlos0_env

import (
	"fmt"
	"os"
	"strings"

	"github.com/caarlos0/env/v9"
)

type Config struct {
	HttpPort   int    `env:"HTTP1_PORT" envDefault:"8080"`
	GrpcPort   int    `env:"HTTP2_PORT" envDefault:"9090"`
	DBHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"sa"`
	DBName     string `env:"DB_NAME" envDefault:"base-gin-go"`
	EntDebug   bool   `env:"ENT_DEBUG" envDefault:"false"`
}

const (
	local = iota
	test
	development
	production
)

var shinonomeEnv = local

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func SetEnv() {
	envName := os.Getenv("PJ_ENV")
	switch strings.ToLower(envName) {
	case "local":
		shinonomeEnv = local
	case "development":
		shinonomeEnv = development
	case "production":
		shinonomeEnv = production
	case "test":
		shinonomeEnv = test
	default:
		panic("PJ_ENV unknown: " + envName + " (available env: local development production test)")
	}
	fmt.Println("Running in [" + envName + "] environment")
}

func SetTestEnv() {
	shinonomeEnv = test
}

func IsTest() bool {
	return shinonomeEnv == test
}

func IsDevelopment() bool {
	return shinonomeEnv == development
}

func IsProduction() bool {
	return shinonomeEnv == production
}
