package log

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	tContext "github.com/tryfix/traceable-context"
)

// NewLog creates a new instance of the logger.
func NewLog(options ...Option) Log {
	opts := new(logOptions)
	opts.applyDefault()
	opts.apply(options...)

	switch opts.output {
	case OutJson:
		return &jsonLogImpl{
			logOptions: opts,
			log:        newZerolog(opts),
		}
	default:
		return &textLogImpl{
			logOptions: opts,
			log:        log.New(opts.writer, ``, log.LstdFlags|log.Lmicroseconds),
		}
	}
}

// Deprecated: Do not use with the json logger
// WithPrefix appends the given prefix to the existing prefix.
func WithPrefix(p string, message interface{}) string {
	return fmt.Sprintf(`%s] [%+v`, p, message)
}

// uuidFromContext extracts the uuid from the given context.
//
// When a uuid is not attached to the context a newly generated uuid will be sent.
func uuidFromContext(ctx context.Context) uuid.UUID {
	uid := tContext.FromContext(ctx)
	if uid == uuid.Nil {
		return uuid.New()
	}

	return uid
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
		return zerolog.ErrorLevel
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
