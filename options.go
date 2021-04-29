package log

import (
	"fmt"
	"io"
	"os"
)

type logOptions struct {
	prefix    string
	suffix    string
	colors    bool
	logLevel  Level
	filePath  bool
	fileDepth int
	writer    io.Writer
}

func (lOpts *logOptions) applyDefault() {
	lOpts.fileDepth = 2
	lOpts.colors = true
	lOpts.logLevel = TRACE
	lOpts.filePath = true
	lOpts.writer = os.Stdout
}

func (lOpts *logOptions) copy() *logOptions {
	return &logOptions{
		prefix:    lOpts.prefix,
		suffix:    lOpts.suffix,
		fileDepth: lOpts.fileDepth,
		colors:    lOpts.colors,
		logLevel:  lOpts.logLevel,
		filePath:  lOpts.filePath,
		writer:    lOpts.writer,
	}
}

func (lOpts *logOptions) apply(options ...Option) {
	for _, opt := range options {
		opt(lOpts)
	}
}

type Option func(*logOptions)

func FileDepth(d int) Option {
	return func(opts *logOptions) {
		opts.fileDepth = d
	}
}

func WithStdOut(w io.Writer) Option {
	return func(opts *logOptions) {
		opts.writer = w
	}
}

func WithFilePath(enabled bool) Option {
	return func(opts *logOptions) {
		opts.filePath = enabled
	}
}

func Prefixed(prefix string) Option {
	return func(opts *logOptions) {
		if opts.prefix != `` {
			opts.prefix = fmt.Sprintf(`%s.%s`, opts.prefix, prefix)
			return
		}
		opts.prefix = prefix
	}
}

func WithColors(enabled bool) Option {
	return func(opts *logOptions) {
		opts.colors = enabled
	}
}

func WithLevel(level Level) Option {
	return func(opts *logOptions) {
		opts.logLevel = level
	}
}
