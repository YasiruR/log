package log_test

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/rs/zerolog"
	"github.com/tryfix/log"
)

// BenchmarkZLBaseline is the simplest benchmark copied from zerolog.
// This benchmark will setup the baseline for the minimalist logger implementations.
func BenchmarkZLBaseline(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).
		Level(zerolog.ErrorLevel).
		With().Timestamp().
		Logger()
	var msg interface{} = "message"
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Error().Msgf("%s", msg)
		}
	})
}

// BenchmarkZLBaselineWithCaller is the simplest benchmark with caller logging enabled.
// This benchmark will setup the baseline for logger implementations that logs the file and line of the caller.
func BenchmarkZLBaselineWithCaller(b *testing.B) {
	logger := zerolog.New(ioutil.Discard).
		Level(zerolog.ErrorLevel).
		With().Timestamp().
		CallerWithSkipFrameCount(3).
		Logger()
	var msg interface{} = "message"
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Error().Msgf("%s", msg)
		}
	})
}

// BenchmarkJsonPrint
func BenchmarkJsonPrint(b *testing.B) {
	logger := log.NewLog(
		log.WithStdOut(ioutil.Discard),
		log.WithOutput(log.OutJson),
	).Log(
		log.WithLevel(log.DEBUG),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Print("message")
		}
	})
}

// BenchmarkJsonLoggers run benchmarks on loggers made with different configurations.
func BenchmarkJsonLoggers(b *testing.B) {
	type config struct {
		name string
		cfs  []log.Option
	}

	// NOTE: Fatal level is not used because it terminates the routine
	configs := []config{
		{name: "Trace", cfs: []log.Option{log.WithLevel(log.TRACE)}},
		{name: "Debug", cfs: []log.Option{log.WithLevel(log.DEBUG)}},
		{name: "Info", cfs: []log.Option{log.WithLevel(log.INFO)}},
		{name: "Warn", cfs: []log.Option{log.WithLevel(log.WARN)}},
		{name: "Error", cfs: []log.Option{log.WithLevel(log.ERROR)}},

		{name: "TracePrefixed", cfs: []log.Option{log.WithLevel(log.TRACE), log.Prefixed("prefix")}},
		{name: "DebugPrefixed", cfs: []log.Option{log.WithLevel(log.DEBUG), log.Prefixed("prefix")}},
		{name: "InfoPrefixed", cfs: []log.Option{log.WithLevel(log.INFO), log.Prefixed("prefix")}},
		{name: "WarnPrefixed", cfs: []log.Option{log.WithLevel(log.WARN), log.Prefixed("prefix")}},
		{name: "ErrorPrefixed", cfs: []log.Option{log.WithLevel(log.ERROR), log.Prefixed("prefix")}},

		{name: "TraceFilepath", cfs: []log.Option{log.WithLevel(log.TRACE), log.WithFilePath(true)}},
		{name: "DebugFilepath", cfs: []log.Option{log.WithLevel(log.DEBUG), log.WithFilePath(true)}},
		{name: "InfoFilepath", cfs: []log.Option{log.WithLevel(log.INFO), log.WithFilePath(true)}},
		{name: "WarnFilepath", cfs: []log.Option{log.WithLevel(log.WARN), log.WithFilePath(true)}},
		{name: "ErrorFilepath", cfs: []log.Option{log.WithLevel(log.ERROR), log.WithFilePath(true)}},
	}

	baseLogger := log.NewLog(log.WithStdOut(ioutil.Discard), log.WithOutput(log.OutJson), log.WithColors(false))

	for _, c := range configs {
		b.Run(c.name, func(b *testing.B) {
			logger := baseLogger.Log(c.cfs...)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					logger.Error("message")
				}
			})
		})
	}
}

func BenchmarkJsonLogInfo(b *testing.B) {
	lg := log.NewLog(
		log.WithLevel(log.INFO),
		log.WithOutput(log.OutJson),
		log.WithStdOut(ioutil.Discard),
		log.WithFilePath(false),
		log.WithColors(false)).Log()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Info(testLog)
		}
	})
}

func BenchmarkJsonLogInfoFilePath(b *testing.B) {
	lg := log.NewLog(
		log.WithLevel(log.INFO),
		log.WithOutput(log.OutJson),
		log.WithStdOut(ioutil.Discard),
		log.WithFilePath(true),
		log.WithColors(false)).Log()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Info(testLog)
		}
	})
}

func BenchmarkJsonInfoContext(b *testing.B) {
	lg := log.NewLog(
		log.WithLevel(log.INFO),
		log.WithOutput(log.OutJson),
		log.WithStdOut(ioutil.Discard),
		log.WithFilePath(false),
		log.WithColors(false)).Log()
	ctx := context.Background()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.InfoContext(ctx, testLog)
		}
	})
}

func BenchmarkJsonInfoContextExt(b *testing.B) {
	ctx1 := context.WithValue(context.Background(), `ctx1`, `ctx one value`)
	ctx2 := context.WithValue(ctx1, `ctx2`, `ctx two value`)
	lg := log.NewLog(
		log.WithLevel(log.INFO),
		log.WithStdOut(ioutil.Discard),
		log.WithFilePath(false),
		log.WithOutput(log.OutJson),
		log.WithCtxExtractor(func(ctx context.Context) []interface{} {
			return []interface{}{ctx.Value(`ctx1`), ctx.Value(`ctx2`)}
		}),
		log.WithColors(false)).Log()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.InfoContext(ctx2, testLog)
		}
	})
}

func BenchmarkJsonInfoParams(b *testing.B) {
	lg := log.NewLog(
		log.WithLevel(log.INFO),
		log.WithOutput(log.OutJson),
		log.WithStdOut(ioutil.Discard),
		log.WithFilePath(false),
		log.WithColors(false)).Log()
	ctx := context.Background()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.InfoContext(ctx, testLog, `parm1`, `parm2`, `parm3`, `parm4`)
		}
	})
}
