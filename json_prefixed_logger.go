package log

import (
	"context"
)

type jsonPrefixedLogger struct {
	jsonLogParser
}

// NewLog creates a new Logger instance using existing config from the PrefixedLogger.
//
// Default configuration values can be overridden by providing Options to the function.
func (l *jsonPrefixedLogger) NewLog(opts ...Option) Logger {
	defaults := l.logOptions.copy()
	defaults.apply(opts...)

	return &jsonLogger{
		jsonLogParser: jsonLogParser{
			logOptions: defaults,
			log:        l.log,
		},
	}
}

// NewPrefixedLog creates a new NewPrefixedLogger instance using existing config from the PrefixedLogger.
//
// Default configuration values can be overridden by providing Options to the function.
func (l *jsonPrefixedLogger) NewPrefixedLog(opts ...Option) PrefixedLogger {
	defaults := l.logOptions.copy()
	defaults.apply(opts...)

	return &jsonPrefixedLogger{
		jsonLogParser: jsonLogParser{
			logOptions: defaults,
			log:        l.log,
		},
	}
}

// Error logs with ERROR level with prefix.
func (l *jsonPrefixedLogger) Error(prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(context.Background(), l.jsonLogParser.log.Error(), prefix, params...).Msgf("%s", message)
}

// Warn logs with WARN level with prefix.
func (l *jsonPrefixedLogger) Warn(prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(context.Background(), l.jsonLogParser.log.Warn(), prefix, params...).Msgf("%s", message)
}

// Info logs with INFO level with prefix.
func (l *jsonPrefixedLogger) Info(prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(context.Background(), l.jsonLogParser.log.Info(), prefix, params...).Msgf("%s", message)
}

// Debug logs with DEBUG level with prefix.
func (l *jsonPrefixedLogger) Debug(prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(context.Background(), l.jsonLogParser.log.Debug(), prefix, params...).Msgf("%s", message)
}

// Trace logs with TRACE level with prefix.
func (l *jsonPrefixedLogger) Trace(prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(context.Background(), l.jsonLogParser.log.Trace(), prefix, params...).Msgf("%s", message)
}

// Fatal logs with FATAL level with prefix.
func (l *jsonPrefixedLogger) Fatal(prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(context.Background(), l.jsonLogParser.log.Fatal(), prefix, params...).Msgf("%s", message)
}

// Fatalln logs with FATAL level with prefix.
func (l *jsonPrefixedLogger) Fatalln(prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(context.Background(), l.jsonLogParser.log.Fatal(), prefix, params...).Msgf("%s", message)
}

// ErrorContext logs with ERROR level with context and prefix.
func (l *jsonPrefixedLogger) ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(ctx, l.jsonLogParser.log.Error(), prefix, params...).Msgf("%s", message)
}

// WarnContext logs with WARN level with context and prefix.
func (l *jsonPrefixedLogger) WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(ctx, l.jsonLogParser.log.Warn(), prefix, params...).Msgf("%s", message)
}

// InfoContext logs with INFO level with context and prefix.
func (l *jsonPrefixedLogger) InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(ctx, l.jsonLogParser.log.Info(), prefix, params...).Msgf("%s", message)
}

// DebugContext logs with DEBUG level with context and prefix.
func (l *jsonPrefixedLogger) DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(ctx, l.jsonLogParser.log.Debug(), prefix, params...).Msgf("%s", message)
}

// TraceContext logs with TRACE level with context and prefix.
func (l *jsonPrefixedLogger) TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(ctx, l.jsonLogParser.log.Trace(), prefix, params...).Msgf("%s", message)
}

// FatalContext logs with FATAL level with context and prefix.
func (l *jsonPrefixedLogger) FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{}) {
	l.jsonLogParser.parse(ctx, l.jsonLogParser.log.Fatal(), prefix, params...).Msgf("%s", message)
}

// Print logs with INFO level.
func (l *jsonPrefixedLogger) Print(v ...interface{}) {
	l.jsonLogParser.log.Print(v...)
}

// Printf logs with INFO level.
func (l *jsonPrefixedLogger) Printf(format string, v ...interface{}) {
	l.jsonLogParser.log.Printf(format, v...)
}

// Println logs with INFO level.
func (l *jsonPrefixedLogger) Println(v ...interface{}) {
	l.jsonLogParser.log.Print(v...)
}
