package log_test

import (
	"io/ioutil"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/tryfix/log"
	traceable_context "github.com/tryfix/traceable-context"
)

var testCtx = traceable_context.WithUUID(uuid.New())

// BenchmarkZLBaseline is the simplest benchmark copied from zerolog.
// This benchmark will setup the baseline for rest of the benchmarks in this file.
func BenchmarkZLBaseline(b *testing.B) {
	logger := zerolog.New(ioutil.Discard)
	var msg interface{} = "message"
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msgf("%s", msg)
		}
	})
}

// BenchmarkJsonLoggers run benchmarks on loggers made with different configurations.
func BenchmarkJsonLoggers(b *testing.B) {
	type config struct {
		name string
		cfs  []log.Option
	}

	configs := []config{
		{name: "Trace", cfs: []log.Option{log.WithLevel(log.TRACE)}},
		// {name: "Debug", cfs: []log.Option{log.WithLevel(log.DEBUG)}},
		// {name: "Info", cfs: []log.Option{log.WithLevel(log.INFO)}},
		// {name: "Warn", cfs: []log.Option{log.WithLevel(log.WARN)}},
		// {name: "Error", cfs: []log.Option{log.WithLevel(log.ERROR)}},
		// {name: "Fatal", cfs: []log.Option{log.WithLevel(log.FATAL)}},
		// {name: "TraceFilepath", cfs: []log.Option{log.WithLevel(log.TRACE), log.WithFilePath(true)}},
		// {name: "DebugFilepath", cfs: []log.Option{log.WithLevel(log.DEBUG), log.WithFilePath(true)}},
		// {name: "InfoFilepath", cfs: []log.Option{log.WithLevel(log.INFO), log.WithFilePath(true)}},
		// {name: "WarnFilepath", cfs: []log.Option{log.WithLevel(log.WARN), log.WithFilePath(true)}},
		// {name: "ErrorFilepath", cfs: []log.Option{log.WithLevel(log.ERROR), log.WithFilePath(true)}},
		// {name: "FatalFilepath", cfs: []log.Option{log.WithLevel(log.FATAL), log.WithFilePath(true)}},
	}

	baseLogger := log.NewLog(log.WithStdOut(ioutil.Discard), log.WithOutput(log.OutJson))

	for _, c := range configs {
		b.Run(c.name, func(b *testing.B) {
			logger := baseLogger.Log(c.cfs...)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					logger.Print("message")
				}
			})
		})
	}
}

func BenchmarkJsonPrint(b *testing.B) {
	logger := log.NewLog(
		log.WithStdOut(ioutil.Discard),
		log.WithOutput(log.OutJson),
	).Log(
		log.WithLevel(log.INFO),
		log.WithFilePath(true),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("message")
		}
	})
}
