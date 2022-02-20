package env_config

import (
	"context"

	"go.uber.org/zap"
)

type Config interface {
	GetInt(key string) int64
	GetString(key string) string
	GetBool(key string) bool
	GetFloat(key string) float64
	GetBinary(key string) []byte
	GetArray(key string) []string
	GetMap(key string) map[string]string
}

// Logger Interface. All methods SHOULD be safe for concurrent use.
type Logger interface {
	// Info logs a message at Info level
	Info(ctx context.Context, msg string, fields ...zap.Field)
	// Warn logs a message at Warn level
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	// Error logs a message at Errors level
	Error(ctx context.Context, msg string, err error)
}
