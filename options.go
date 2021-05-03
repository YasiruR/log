package log

import (
	"fmt"
	"io"
	"os"
)

// Option represents a function that does one or more alterations to 'logOptions' inside the logger.
type Option func(*logOptions)

// logOptions contains all the configuration options for the logger.
type logOptions struct {
	prefix    string
	suffix    string
	colors    bool
	logLevel  Level
	filePath  bool
	fileDepth int
	ctxKeys   []interface{}
	writer    io.Writer
}

// applyDefault applies a set of predefined configurations to the logger.
func (lOpts *logOptions) applyDefault() {
	lOpts.fileDepth = 2
	lOpts.colors = true
	lOpts.logLevel = TRACE
	lOpts.filePath = true
	lOpts.writer = os.Stdout
}

// copy returns a copy of existing configuration values of the logger.
func (lOpts *logOptions) copy() *logOptions {
	return &logOptions{
		prefix:    lOpts.prefix,
		suffix:    lOpts.suffix,
		fileDepth: lOpts.fileDepth,
		colors:    lOpts.colors,
		logLevel:  lOpts.logLevel,
		filePath:  lOpts.filePath,
		ctxKeys:   lOpts.ctxKeys,
		writer:    lOpts.writer,
	}
}

// apply applies given configuration values to the logger.
//
// This will replace existing configuration values with provided values.
func (lOpts *logOptions) apply(options ...Option) {
	for _, opt := range options {
		opt(lOpts)
	}
}

// FileDepth
//
// TODO: add description
func FileDepth(d int) Option {
	return func(opts *logOptions) {
		opts.fileDepth = d
	}
}

// WithStdOut sets the log writer.
func WithStdOut(w io.Writer) Option {
	return func(opts *logOptions) {
		opts.writer = w
	}
}

// WithFilePath sets whether the file path is logged or not.
func WithFilePath(enabled bool) Option {
	return func(opts *logOptions) {
		opts.filePath = enabled
	}
}

// Prefixed appends the given prefix value to the existing prefix value.
func Prefixed(prefix string) Option {
	return func(opts *logOptions) {
		if opts.prefix != `` {
			opts.prefix = fmt.Sprintf(`%s.%s`, opts.prefix, prefix)
			return
		}
		opts.prefix = prefix
	}
}

// WithColors enables colours in log messages.
func WithColors(enabled bool) Option {
	return func(opts *logOptions) {
		opts.colors = enabled
	}
}

// WithLevel sets the log level.
//
// The log level is used to determine which types of logs are logged depending on the precedence of the log level.
// TODO: need a clearer explanation
func WithLevel(level Level) Option {
	return func(opts *logOptions) {
		opts.logLevel = level
	}
}

// WithCtxKeys sets a set of context keys to be used to extract values from the context and add them to the log entry.
func WithCtxKeys(keys ...interface{}) Option {
	return func(opts *logOptions) {
		// don't proceed when no new keys are provided
		if len(keys) == 0 {
			return
		}

		// remove duplicates
		allKeys := append(opts.ctxKeys, keys...)
		m := make(map[interface{}]bool)
		for _, k := range allKeys {
			m[k] = true
		}

		var uniqueKeys []interface{}
		for mk := range m {
			uniqueKeys = append(uniqueKeys, mk)
		}

		opts.ctxKeys = uniqueKeys
	}
}
