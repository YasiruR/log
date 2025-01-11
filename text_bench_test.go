package log_test

import (
	"context"
	"io/ioutil"
	nativeLog "log"
	"testing"

	"github.com/tryfix/log"
)

const testLog = `test log entry, test log entry, test log entry, test log entry, test log entry`

func BenchmarkGoNative(b *testing.B) {
	lg := nativeLog.New(ioutil.Discard, `test-logger.path`, nativeLog.Lmicroseconds)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Println(testLog)
		}
	})
}

func BenchmarkTextLogInfo(b *testing.B) {
	lg := log.NewLog(
		log.WithLevel(log.INFO),
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

func BenchmarkTextLogInfoFilePath(b *testing.B) {
	lg := log.NewLog(
		log.WithLevel(log.INFO),
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

func BenchmarkTextInfoContext(b *testing.B) {
	lg := log.NewLog(
		log.WithLevel(log.INFO),
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

func BenchmarkTextInfoContextExt(b *testing.B) {
	ctx1 := context.WithValue(context.Background(), `ctx1`, `ctx one value`)
	ctx2 := context.WithValue(ctx1, `ctx2`, `ctx two value`)
	lg := log.NewLog(
		log.WithLevel(log.INFO),
		log.WithStdOut(ioutil.Discard),
		log.WithFilePath(false),
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

func BenchmarkTextInfoParams(b *testing.B) {
	lg := log.NewLog(
		log.WithLevel(log.INFO),
		log.WithStdOut(ioutil.Discard),
		log.WithFilePath(false),
		log.WithColors(false)).Log()
	ctx := context.Background()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.InfoContext(ctx, testLog,
				`parm1`, `parm2`, `parm3`, `parm4`)
		}
	})
}
