package config

import (
	"auth-service/repository"
	"database/sql"

	"go.uber.org/zap"
)

type Config struct {
	Db   *sql.DB
	Repo repository.Repository
	Log  *zap.Logger
}

func NewConfig(conn *sql.DB, repo repository.Repository, log *zap.Logger) *Config {
	return &Config{
		Db:   conn,
		Repo: repo,
		Log:  log,
	}
}
