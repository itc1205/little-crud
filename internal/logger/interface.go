package logger

import "context"

type Logger interface {
	Log(ctx context.Context, message LogMessage) error
}
