package log

import "context"

var Constructor = NewLog(FileDepth(2))
var StdLogger = Constructor.Log(FileDepth(3))
var PrefixedStdLogger = Constructor.PrefixedLog(FileDepth(3))

func Fatal(message interface{}, params ...interface{}) {
	StdLogger.Fatal(message, params...)
}

func Error(message interface{}, params ...interface{}) {
	StdLogger.Error(message, params...)
}

func Warn(message interface{}, params ...interface{}) {
	StdLogger.Warn(message, params...)
}

func Debug(message interface{}, params ...interface{}) {
	StdLogger.Debug(message, params...)
}

func Info(message interface{}, params ...interface{}) {
	StdLogger.Info(message, params...)
}

func Trace(message interface{}, params ...interface{}) {
	StdLogger.Trace(message, params...)
}

func FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.FatalContext(ctx, message, params...)
}

func ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.ErrorContext(ctx, message, params...)
}

func WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.WarnContext(ctx, message, params...)
}

func DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.DebugContext(ctx, message, params...)
}

func InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.InfoContext(ctx, message, params...)
}

func TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	StdLogger.TraceContext(ctx, message, params...)
}
