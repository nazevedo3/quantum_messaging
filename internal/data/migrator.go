package data

import (
	"database/sql"
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// NewMigrator retrieves the location of the migration files for the test database
func NewMigrator(db *sql.DB) (*migrate.Migrate, error) {
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("failed to get path")
	}
	sourceUrl := "file://" + filepath.Dir(path) + "../../../migrations"
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create migrator driver: %s", err)
	}
	m, err := migrate.NewWithDatabaseInstance(sourceUrl, "postgres", driver)
	return m, err
}
