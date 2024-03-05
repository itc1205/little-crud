package natsclick

import (
	"context"
	"encoding/json"

	"github.com/itc1205/little-crud/internal/logger"
	"github.com/nats-io/nats.go"
)

type NatsClickLog struct {
	nc   *nats.Conn
	subj string
}

func New(nc *nats.Conn, subj string) *NatsClickLog {
	return &NatsClickLog{nc: nc, subj: subj}
}

func (nl *NatsClickLog) Log(ctx context.Context, message logger.LogMessage) error {
	msg_json, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = nl.nc.Publish(nl.subj, msg_json)
	return err
}
