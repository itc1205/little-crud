package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/itc1205/little-crud/internal/config"
)

func InitSqlDB(cfg config.PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", formatConnect(cfg))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}

func formatConnect(cfg config.PostgresConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PgUser, cfg.PgPwd, cfg.PgHost, cfg.PgPort, cfg.PgDBName,
	)
}
