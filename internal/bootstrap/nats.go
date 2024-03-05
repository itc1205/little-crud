package bootstrap

import (
	"fmt"

	"github.com/itc1205/little-crud/internal/config"
	"github.com/nats-io/nats.go"
)

func InitNatsConn(cfg config.NatsConfig) (*nats.Conn, error) {
	nc, err := nats.Connect(natsUrlFromConfig(cfg))
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func natsUrlFromConfig(cfg config.NatsConfig) string {
	return fmt.Sprintf("%s:%s", cfg.NHost, cfg.NPort)
}
