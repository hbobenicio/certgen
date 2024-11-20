package application

import (
	"certgen/internal/config"
	"certgen/internal/database"
	"database/sql"
	"fmt"
)

type App struct {
	Cfg *config.Config
	DB  *sql.DB
}

func Setup() (*App, error) {

	app := &App{}

	app.Cfg = config.New()
	if err := app.Cfg.LoadFromEnv(); err != nil {
		return nil, fmt.Errorf("config loading failed: %w", err)
	}

	db, err := database.NewPool(app.Cfg)
	if err != nil {
		return nil, fmt.Errorf("db loading failed: %w", err)
	}
	app.DB = db

	if err := database.Migrate(app.Cfg, db); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return app, nil
}

func (app *App) Cleanup() {
	app.DB.Close()
}
