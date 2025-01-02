package database

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/umed-hotamov/url-shortener/internal/config"
	"go.uber.org/zap"
)

func InitDB(cfg *config.Config, logger *zap.Logger) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		logger.Fatal("Failed to connect to postgres: %s", zap.String("dsn string", dsn))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal("Failed to ping postgres: %s", zap.String("dsn string", dsn))
		return nil, err
	}

	return db, nil
}
