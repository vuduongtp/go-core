package logger

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
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

// LogPanic log panic
func LogPanic(content ...interface{}) {
	logadapter.Panic(content...)
}

// LogErrorf for logging errors with context to log request_id and correlation_id
func LogErrorf(ctx context.Context, format string, a ...any) {
	logadapter.LogWithContext(ctx, fmt.Sprintf(format, a...), logadapter.LogTypeError)
}

// LogDebugf for logging debug with context to log request_id and correlation_id
func LogDebugf(ctx context.Context, format string, a ...any) {
	logadapter.LogWithContext(ctx, fmt.Sprintf(format, a...), logadapter.LogTypeDebug)
}

// LogInfof for logging info with context to log request_id and correlation_id
func LogInfof(ctx context.Context, format string, a ...any) {
	logadapter.LogWithContext(ctx, fmt.Sprintf(format, a...), logadapter.LogTypeInfo)
}

// LogWarnf for logging warn with context to log request_id and correlation_id
func LogWarnf(ctx context.Context, format string, a ...any) {
	logadapter.LogWithContext(ctx, fmt.Sprintf(format, a...), logadapter.LogTypeWarn)
}

// LogWithContext log content with context
// content[0] : message -> interface{},
// content[1] : log type -> string,
// content[2] : log field -> map[string]interface{}
func LogWithContext(ctx context.Context, content ...interface{}) {
	logadapter.LogWithContext(ctx, content...)
}

// Ctx get logrus entry from context
func Ctx(ctx context.Context) *logrus.Entry {
	return logadapter.SetContext(ctx)
}

// AddLogField add more log field to context
func AddLogField(ctx context.Context, key string, value interface{}) context.Context {
	return logadapter.SetCustomLogField(ctx, key, value)
}
