package log

import (
	"context"

	"github.com/rs/zerolog"
)

// logParser contains parsing logic for a logger.
type jsonLogParser struct {
	*logOptions
	log zerolog.Logger
}

// parse parses all additional data.
func (l *jsonLogParser) parse(ctx context.Context, event *zerolog.Event, prefix string, params ...interface{}) *zerolog.Event {
	event = l.withUUID(ctx, event)
	event = l.withPrefix(event, prefix)
	event = l.withExtractedCtx(ctx, event)
	event = l.withParams(event, params...)

	return event
}

// withLoggerPrefix attaches the logger prefix to the event.
func (l *jsonLogParser) withPrefix(event *zerolog.Event, prefix string) *zerolog.Event {
	key := "prefix"

	if l.prefix != "" {
		if prefix != "" {
			return event.Str(key, l.prefix+"."+prefix)
		}

		return event.Str(key, l.prefix)
	}

	if prefix != "" {
		return event.Str(key, prefix)
	}

	return event
}

// withUUID attaches the uuid from context to event.
func (l *jsonLogParser) withUUID(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	return event.Str("uuid", uuidFromContext(ctx).String())
}

// withExtractedCtx adds the extacted context values to the event.
func (l *jsonLogParser) withExtractedCtx(ctx context.Context, event *zerolog.Event) *zerolog.Event {
	if l.ctxExt != nil {
		if ctxData := l.ctxExt(ctx); len(ctxData) > 0 {
			return event.Interface("context", ctxData)
		}
	}

	return event
}

// withParams adds additional details to the event.
func (l *jsonLogParser) withParams(event *zerolog.Event, params ...interface{}) *zerolog.Event {
	if len(params) == 0 {
		return event
	}

	return event.Interface("details", params)
}
