package log

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	tContext "github.com/tryfix/traceable-context"
)

// NewLog creates a new instance of the logger.
func NewLog(options ...Option) Log {
	opts := new(logOptions)
	opts.applyDefault()
	opts.apply(options...)

	switch opts.output {
	case OutJson:
		return &textLogImpl{
			logOptions: opts,
			log:        log.New(opts.writer, ``, log.LstdFlags|log.Lmicroseconds),
		}
	default:
		return &textLogImpl{
			logOptions: opts,
			log:        log.New(opts.writer, ``, log.LstdFlags|log.Lmicroseconds),
		}
	}
}

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
