package migrate

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func postgresGetMigrator(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://migrations/postgres", "postgres", driver)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func PostgresMigrateUp(db *sql.DB) error {
	m, err := postgresGetMigrator(db)
	if err != nil {
		return err
	}
	err = m.Up()
	return err
}

func PostgresMigrateDown(db *sql.DB) error {
	m, err := postgresGetMigrator(db)
	if err != nil {
		return err
	}
	err = m.Up()
	return err
}
