package database

import (
	"certgen/internal/config"
	"database/sql"
	"fmt"
)

func NewPool(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DbUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to open/connect to database: %w", err)
	}

	return db, nil
}

func Migrate(cfg *config.Config, db *sql.DB) error {
	sql, err := GetCreateTableCertsSql()
	if err != nil {
		return err
	}

	if _, err := db.Exec(sql); err != nil {
		return err
	}

	return nil
}
