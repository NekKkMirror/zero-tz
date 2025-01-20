package db

import (
	"database/sql"
	"fmt"
	"github.com/NekKkMirror/zero-tz/internal/pkg/config"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

func InitDB(cfg *config.Config, log *logrus.Logger) (*reform.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	sqlDB, err := sql.Open(postgresql.Dialect.String(), dsn)
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	reformDB := reform.NewDB(sqlDB, postgresql.Dialect, nil)

	log.Info("DB connection established.")

	return reformDB, nil
}
