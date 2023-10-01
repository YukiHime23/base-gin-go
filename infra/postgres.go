package infra

import (
	config_caarlos0_env "base-gin-go/config/caarlos0-env"
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresOpen(ctx context.Context, cfg *config_caarlos0_env.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%v sslmode=disable",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		err = fmt.Errorf("[infra.PostgresOpen] failed to open connection to database: %w", err)
	}

	return
}
