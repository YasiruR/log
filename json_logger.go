package log

import (
	"context"
)

type jsonLogger struct {
	jsonLogParser
}

// NewLog creates a new Logger instance using existing config from the Logger.
//
// Default configuration values can be overridden by providing Options to the function.
func (l *jsonLogger) NewLog(opts ...Option) Logger {
	defaults := l.logOptions.copy()
	defaults.apply(opts...)

	return &jsonLogger{
		jsonLogParser: jsonLogParser{
			logOptions: defaults,
			log:        l.log,
		},
	}
}

// NewPrefixedLog creates a new NewPrefixedLogger instance using existing config from the Logger.
//
// Default configuration values can be overridden by providing Options to the function.
func (l *jsonLogger) NewPrefixedLog(opts ...Option) PrefixedLogger {
	defaults := l.logOptions.copy()
	defaults.apply(opts...)

	return &jsonPrefixedLogger{
		jsonLogParser: jsonLogParser{
			logOptions: defaults,
			log:        l.log,
		},
	}
}

// Error logs with ERROR level.
func (l *jsonLogger) Error(message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Error().Msgf("%s", message)
}

// Warn logs with WARN level.
func (l *jsonLogger) Warn(message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Warn().Msgf("%s", message)
}

// Info logs with INFO level.
func (l *jsonLogger) Info(message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Info().Msgf("%s", message)
}

// Debug logs with DEBUG level.
func (l *jsonLogger) Debug(message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Debug().Msgf("%s", message)
}

// Trace logs with TRACE level.
func (l *jsonLogger) Trace(message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Trace().Msgf("%s", message)
}

// Fatal logs with FATAL level.
func (l *jsonLogger) Fatal(message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Fatal().Msgf("%s", message)
}

// Fatalln logs with FATAL level.
func (l *jsonLogger) Fatalln(message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Fatal().Msgf("%s", message)
}

// ErrorContext logs with ERROR level with context.
func (l *jsonLogger) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Error().Msgf("%s", message)
}

// WarnContext logs with WARN level with context.
func (l *jsonLogger) WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Warn().Msgf("%s", message)
}

// InfoContext logs with INFO level with context.
func (l *jsonLogger) InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Info().Msgf("%s", message)
}

// DebugContext logs with DEBUG level with context.
func (l *jsonLogger) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Debug().Msgf("%s", message)
}

// TraceContext logs with TRACE level with context.
func (l *jsonLogger) TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Trace().Msgf("%s", message)
}

// FatalContext logs with FATAL level with context.
func (l *jsonLogger) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.jsonLogParser.log.Fatal().Msgf("%s", message)
}

// Print logs with INFO level.
func (l *jsonLogger) Print(v ...interface{}) {
	l.jsonLogParser.log.Print(v...)
}

// Printf logs with INFO level.
func (l *jsonLogger) Printf(format string, v ...interface{}) {
	l.jsonLogParser.log.Printf(format, v...)
}

// Println logs with INFO level.
func (l *jsonLogger) Println(v ...interface{}) {
	l.jsonLogParser.log.Print(v...)
}
