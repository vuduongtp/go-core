package logger

import (
	"context"

	"github.com/vuduongtp/go-logadapter"
)

// LogError for logging errors with context to log request_id and correlation_id
func LogError(ctx context.Context, content interface{}) {
	logadapter.LogWithContext(ctx, content, logadapter.LogTypeError)
}

// LogDebug for logging debug with context to log request_id and correlation_id
func LogDebug(ctx context.Context, content interface{}) {
	logadapter.LogWithContext(ctx, content, logadapter.LogTypeDebug)
}

// LogInfo for logging info with context to log request_id and correlation_id
func LogInfo(ctx context.Context, content interface{}) {
	logadapter.LogWithContext(ctx, content, logadapter.LogTypeInfo)
}

// LogWarn for logging warn with context to log request_id and correlation_id
func LogWarn(ctx context.Context, content interface{}) {
	logadapter.LogWithContext(ctx, content, logadapter.LogTypeWarn)
}

// LogWithContext log content with context
// content[0] : message -> interface{},
// content[1] : log type -> string,
// content[2] : log field -> map[string]interface{}
func LogWithContext(ctx context.Context, content ...interface{}) {
	logadapter.LogWithContext(ctx, content...)
}
