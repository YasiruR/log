package log

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/google/uuid"
	tContext "github.com/tryfix/traceable-context"
)

// logMessage represents the overall message being logged.
type logMessage struct {
	typ     string
	color   string
	message interface{}
	uuid    string
	file    string
	line    int
}

// logParser contains parsing logic for a logger.
type logParser struct {
	*logOptions
	log *log.Logger
}

//isLoggable checks whether it is possible to log in the given level under current configurations.
func (l *logParser) isLoggable(level Level) bool {
	return logTypes[level] <= logTypes[l.logLevel]
}

// colored colour encodes the log level tag.
//
// Whether this returns coloured tags or not depends on the colour configuration of the logger.
func (l *logParser) colored(level Level) string {
	if l.colors {
		return string(logColors[level])
	}

	return string(level)
}

// WithPrefix appends the given prefix to the existing prefix.
func (l *logParser) WithPrefix(p string, message interface{}) string {
	if l.prefix != `` {
		if p == `` {
			return fmt.Sprintf(`%s] [%+v`, l.prefix, message)
		}

		return fmt.Sprintf(`%s.%s] [%+v`, l.prefix, p, message)
	}

	return fmt.Sprintf(`%s] [%+v`, p, message)
}

// logEntry prints the log entry to the configured io.Writer.
func (l *logParser) logEntry(ctx context.Context, level Level, message interface{}, prms ...interface{}) {
	if !l.isLoggable(level) {
		return
	}

	var params []interface{}
	format := "%s [%s] [%+v]"
	// logLevel := string(level)
	logLevel := l.colored(level)
	uid := uuidFromContext(ctx).String()

	logMsg := &logMessage{
		typ:     logLevel,
		message: message,
		uuid:    uid,
	}

	params = append(params, logLevel, uid, fmt.Sprintf(`%s`, message))

	funcName := ``
	file := `<Unknown>`
	line := 1
	pc, file, line, ok := runtime.Caller(l.fileDepth)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
	}

	format = "%s [%s] [%+v" + fmt.Sprintf(` on func %s`, funcName) + "]"

	if l.filePath {
		logMsg.file = file
		logMsg.line = line
		format = "%s [%s] [%+v" + fmt.Sprintf(` on func %s on %s line %d`, funcName, file, line) + "]"
	}

	if len(prms) > 0 {
		format += " %+v"
		params = append(params, prms)
	}

	// add context details
	if l.ctxExt != nil {
		if ctxData := l.ctxExt(ctx); len(ctxData) > 0 {
			format += " %v"
			params = append(params, ctxData)
		}
	}

	if level == FATAL {
		l.log.Fatalf(format, params...)
	}

	l.log.Printf(format, params...)
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

// // getFromCtx extracts values from the context that has keys defined in ctxKeys.
// func getFromCtx(ctx context.Context, keys []interface{}) string {
// 	var s string

// 	for _, k := range keys {
// 		v := ctx.Value(k)
// 		if v != nil {
// 			s += ", " + fmt.Sprintf("%s: %+v", k, v)
// 		}
// 	}

// 	return s[2:]
// }
