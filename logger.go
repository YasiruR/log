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

type PrefixedLogger interface {
	Fatal(prefix string, message interface{}, params ...interface{})
	Error(prefix string, message interface{}, params ...interface{})
	Warn(prefix string, message interface{}, params ...interface{})
	Debug(prefix string, message interface{}, params ...interface{})
	Info(prefix string, message interface{}, params ...interface{})
	Trace(prefix string, message interface{}, params ...interface{})
	FatalContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	ErrorContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	WarnContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	DebugContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	InfoContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	TraceContext(ctx context.Context, prefix string, message interface{}, params ...interface{})
	NewLog(...Option) Logger
	NewPrefixedLog(opts ...Option) PrefixedLogger
	SimpleLogger
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

func (l *logger) ErrorContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(ERROR, ctx, l.WithPrefix(``, message), params...)
}

func (l *logger) WarnContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(WARN, ctx, l.WithPrefix(``, message), params...)
}

func (l *logger) InfoContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(INFO, ctx, l.WithPrefix(``, message), params...)
}

func (l *logger) DebugContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(DEBUG, ctx, l.WithPrefix(``, message), params...)
}

func (l *logger) TraceContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(TRACE, ctx, l.WithPrefix(``, message), params...)
}

func (l *logger) Error(message interface{}, params ...interface{}) {
	l.logEntry(ERROR, nil, l.WithPrefix(``, message), params...)
}

func (l *logger) Warn(message interface{}, params ...interface{}) {
	l.logEntry(WARN, nil, l.WithPrefix(``, message), params...)
}

func (l *logger) Info(message interface{}, params ...interface{}) {
	l.logEntry(INFO, nil, l.WithPrefix(``, message), params...)
}

func (l *logger) Debug(message interface{}, params ...interface{}) {
	l.logEntry(DEBUG, nil, l.WithPrefix(``, message), params...)
}

func (l *logger) Trace(message interface{}, params ...interface{}) {
	l.logEntry(TRACE, nil, l.WithPrefix(``, message), params...)
}

func (l *logger) Fatal(message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, l.WithPrefix(``, message), params...)
}

func (l *logger) Fatalln(message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, l.WithPrefix(``, message), params...)
}

func (l *logger) FatalContext(ctx context.Context, message interface{}, params ...interface{}) {
	l.logEntry(FATAL, nil, message, params)
}

func (l *logger) Print(v ...interface{}) {
	l.logEntry(INFO, nil, l.WithPrefix(``, v), `INFO`)
}

func (l *logger) Printf(format string, v ...interface{}) {
	l.logEntry(INFO, nil, l.WithPrefix(``, fmt.Sprintf(format, v...)), `INFO`)
}

func (l *logger) Println(v ...interface{}) {
	l.logEntry(INFO, nil, l.WithPrefix(``, v), `INFO`)
}
