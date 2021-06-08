package log

import (
	"github.com/rs/zerolog"
)

type jsonLogImpl struct {
	*logOptions
	log zerolog.Logger
}

func newJsonLogImpl(opts *logOptions) Log {
	return &jsonLogImpl{
		logOptions: opts,
		log:        newZerolog(opts),
	}
}

func (l *jsonLogImpl) Log(options ...Option) Logger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return newJsonLogger(opts)
}

func (*jsonLogImpl) SimpleLog() SimpleLogger {
	panic(`implement me`)
}

func (l *jsonLogImpl) PrefixedLog(options ...Option) PrefixedLogger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return newPrefixedJsonLogger(opts)
}

// zerologLevel convers the config log level to corresponding zerolog log level.
func zerologLevel(lvl Level) zerolog.Level {
	switch lvl {
	case FATAL:
		return zerolog.FatalLevel
	case ERROR:
		return zerolog.ErrorLevel
	case WARN:
		return zerolog.WarnLevel
	case INFO:
		return zerolog.InfoLevel
	case DEBUG:
		return zerolog.DebugLevel
	case TRACE:
		return zerolog.TraceLevel
	default:
		return zerolog.TraceLevel
	}
}

// newZerolog creates a new zerolog instance using given configs.
func newZerolog(opts *logOptions) zerolog.Logger {
	z := zerolog.New(opts.writer).With().Timestamp()

	if opts.filePath {
		z = z.CallerWithSkipFrameCount(opts.fileDepth)
	}

	return z.Logger().Level(zerologLevel(opts.logLevel))
}
