package architecture

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDB(log *slog.Logger, config Config) (*sqlx.DB, error) {
	log.Info("Connecting to database")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s search_path=%s port=%s sslmode=disable", config.Database.User, config.Database.Password, config.Database.Database, config.Database.Host, config.Database.Schema, config.Database.Port))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Info("schema name set successfully", "name", config.Database.Schema)
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	log.Info("Database connection successful")
	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("%s", config.Database.MigrationSrc), "postgres", driver)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}
	log.Info("Database migrations completed")
	return db, nil
}
