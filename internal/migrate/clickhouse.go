package migrate

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/clickhouse"
)

func clickGetMigrator(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := clickhouse.WithInstance(db, &clickhouse.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations/clickhouse",
		"clickhouse", driver,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func ClickMigrateUp(db *sql.DB) error {
	m, err := clickGetMigrator(db)
	if err != nil {
		return err
	}
	return m.Up()
}

func ClickMigrateDown(db *sql.DB) error {
	m, err := clickGetMigrator(db)
	if err != nil {
		return err
	}
	return m.Down()
}
