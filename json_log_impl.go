package log

import (
	"github.com/rs/zerolog"
)

type jsonLogImpl struct {
	*logOptions
	log zerolog.Logger
}

func (l *jsonLogImpl) Log(options ...Option) Logger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return &jsonLogger{
		jsonLogParser: jsonLogParser{
			logOptions: opts,
			log:        l.log,
		},
	}
}

func (*jsonLogImpl) SimpleLog() SimpleLogger {
	panic(`implement me`)
}

func (l *jsonLogImpl) PrefixedLog(options ...Option) PrefixedLogger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return &jsonPrefixedLogger{
		jsonLogParser: jsonLogParser{
			logOptions: opts,
			log:        l.log,
		},
	}
}
