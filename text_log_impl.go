package log

import "log"

type textLogImpl struct {
	*logOptions
	log *log.Logger
}

func (l *textLogImpl) Log(options ...Option) Logger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return &logger{
		logParser: logParser{
			logOptions: opts,
			log:        l.log,
		},
	}
}

func (*textLogImpl) SimpleLog() SimpleLogger {
	panic(`implement me`)
}

func (l *textLogImpl) PrefixedLog(options ...Option) PrefixedLogger {
	opts := l.logOptions.copy()
	opts.apply(options...)

	return &prefixedLogger{
		logParser: logParser{
			logOptions: opts,
			log:        l.log,
		},
	}
}
