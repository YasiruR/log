package log

import (
	"context"
	"fmt"
	"log"
)

type SimpleLogger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type Logger interface {
	Fatal(message interface{}, params ...interface{})
	Error(message interface{}, params ...interface{})
	Warn(message interface{}, params ...interface{})
	Debug(message interface{}, params ...interface{})
	Info(message interface{}, params ...interface{})
	Trace(message interface{}, params ...interface{})
	FatalContext(ctx context.Context, message interface{}, params ...interface{})
	ErrorContext(ctx context.Context, message interface{}, params ...interface{})
	WarnContext(ctx context.Context, message interface{}, params ...interface{})
	DebugContext(ctx context.Context, message interface{}, params ...interface{})
	InfoContext(ctx context.Context, message interface{}, params ...interface{})
	TraceContext(ctx context.Context, message interface{}, params ...interface{})
	SimpleLogger
	NewLog(...Option) Logger
	NewPrefixedLog(opts ...Option) PrefixedLogger
}

type Log interface {
	Log(...Option) Logger
	SimpleLog() SimpleLogger
	PrefixedLog(...Option) PrefixedLogger
}

type logIpml struct {
	log *log.Logger
	*logOptions
}

func NewLog(options ...Option) Log {
	opts := new(logOptions)
	opts.applyDefault()
	opts.apply(options...)

	return &logIpml{
		log:        log.New(opts.writer, ``, log.LstdFlags|log.Lmicroseconds),
		logOptions: opts,
	}
}

func (l *logIpml) Log(options ...Option) Logger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return &logger{
		logParser: logParser{
			logOptions: opts,
			log:        l.log,
		},
	}
}

func (*logIpml) SimpleLog() SimpleLogger {
	panic(`implement me`)
}

func (l *logIpml) PrefixedLog(options ...Option) PrefixedLogger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return &prefixedLogger{
		logParser: logParser{
			logOptions: opts,
			log:        l.log,
		},
	}
}

type logger struct {
	logParser
}

// NewLog creates a new Logger instance using existing config from the Logger.
//
// Default configuration values can be overridden by providing Options to the function.
func (l *logger) NewLog(opts ...Option) Logger {
	defaults := l.logOptions.copy()
	defaults.apply(opts...)

	return &logger{
		logParser: logParser{
			logOptions: defaults,
			log:        l.log,
		},
	}
}

// NewPrefixedLog creates a new NewPrefixedLogger instance using existing config from the Logger.
//
// Default configuration values can be overridden by providing Options to the function.
func (l *logger) NewPrefixedLog(opts ...Option) PrefixedLogger {
	defaults := l.logOptions.copy()
	defaults.apply(opts...)

	return &prefixedLogger{
		logParser: logParser{
			logOptions: defaults,
			log:        l.log,
		},
	}
}

// ErrorContext logs with ERROR level with context.
func (l *logger) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(ctx, ERROR, l.WithPrefix(``, message), params...)
}

// WarnContext logs with WARN level with context.
func (l *logger) WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(ctx, WARN, l.WithPrefix(``, message), params...)
}

// InfoContext logs with INFO level with context.
func (l *logger) InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(ctx, INFO, l.WithPrefix(``, message), params...)
}

// DebugContext logs with DEBUG level with context.
func (l *logger) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(ctx, DEBUG, l.WithPrefix(``, message), params...)
}

// TraceContext logs with TRACE level with context.
func (l *logger) TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(ctx, TRACE, l.WithPrefix(``, message), params...)
}

// FatalContext logs with FATAL level with context.
func (l *logger) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(context.Background(), FATAL, message, params)
}

// Error logs with ERROR level.
func (l *logger) Error(message interface{}, params ...interface{}) {
	l.logEntry(context.Background(), ERROR, l.WithPrefix(``, message), params...)
}

// Warn logs with WARN level.
func (l *logger) Warn(message interface{}, params ...interface{}) {
	l.logEntry(context.Background(), WARN, l.WithPrefix(``, message), params...)
}

// Info logs with INFO level.
func (l *logger) Info(message interface{}, params ...interface{}) {
	l.logEntry(context.Background(), INFO, l.WithPrefix(``, message), params...)
}

// Debug logs with DEBUG level.
func (l *logger) Debug(message interface{}, params ...interface{}) {
	l.logEntry(context.Background(), DEBUG, l.WithPrefix(``, message), params...)
}

// Trace logs with TRACE level.
func (l *logger) Trace(message interface{}, params ...interface{}) {
	l.logEntry(context.Background(), TRACE, l.WithPrefix(``, message), params...)
}

// Fatal logs with FATAL level.
func (l *logger) Fatal(message interface{}, params ...interface{}) {
	l.logEntry(context.Background(), FATAL, l.WithPrefix(``, message), params...)
}

// Fatalln logs with FATAL level.
func (l *logger) Fatalln(message interface{}, params ...interface{}) {
	l.logEntry(context.Background(), FATAL, l.WithPrefix(``, message), params...)
}

// Print logs with INFO level.
func (l *logger) Print(v ...interface{}) {
	l.logEntry(context.Background(), INFO, l.WithPrefix(``, v), `INFO`)
}

// Printf logs with INFO level.
func (l *logger) Printf(format string, v ...interface{}) {
	l.logEntry(context.Background(), INFO, l.WithPrefix(``, fmt.Sprintf(format, v...)), `INFO`)
}

// Println logs with INFO level.
func (l *logger) Println(v ...interface{}) {
	l.logEntry(context.Background(), INFO, l.WithPrefix(``, v), `INFO`)
}
