package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/ClickHouse/clickhouse-go"

	"github.com/itc1205/little-crud/internal/config"
)

func InitClickHouse(cfg config.ClickHouseConfig) (*sql.DB, error) {
	db, err := sql.Open("clickhouse", formatClickConnect(cfg))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func formatClickConnect(cfg config.ClickHouseConfig) string {
	return fmt.Sprintf(
		"clickhouse://%s:%s?username=%s&password=%s&database=%s&x-multi-statement=true",
		cfg.ChHost,
		cfg.ChPort,
		cfg.ChUser,
		cfg.ChPwd,
		cfg.ChDBName,
	)
}
