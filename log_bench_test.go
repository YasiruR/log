package log

import (
	"context"
	"io/ioutil"
	"log"
	"testing"
)

const testLog = `test log entry, test log entry, test log entry, test log entry, test log entry`

func BenchmarkGoNative(b *testing.B) {
	lg := log.New(ioutil.Discard, `test-logger.path`, log.Lmicroseconds)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Println(testLog)
		}
	})
}

func BenchmarkTextLogInfo(b *testing.B) {
	lg := NewLog(
		WithLevel(INFO),
		WithStdOut(ioutil.Discard),
		WithFilePath(false),
		WithColors(false)).Log()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Info(testLog)
		}
	})
}

func BenchmarkTextLogInfoFilePath(b *testing.B) {
	lg := NewLog(
		WithLevel(INFO),
		WithStdOut(ioutil.Discard),
		WithFilePath(true),
		WithColors(false)).Log()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Info(testLog)
		}
	})
}

func BenchmarkTextInfoContext(b *testing.B) {
	lg := NewLog(
		WithLevel(INFO),
		WithStdOut(ioutil.Discard),
		WithFilePath(false),
		WithColors(false)).Log()
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
	lg := NewLog(
		WithLevel(INFO),
		WithStdOut(ioutil.Discard),
		WithFilePath(false),
		WithCtxExtractor(func(ctx context.Context) []interface{} {
			return []interface{}{ctx.Value(`ctx1`), ctx.Value(`ctx2`)}
		}),
		WithColors(false)).Log()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.InfoContext(ctx2, testLog)
		}
	})
}

func BenchmarkTextInfoParams(b *testing.B) {
	lg := NewLog(
		WithLevel(INFO),
		WithStdOut(ioutil.Discard),
		WithFilePath(false),
		WithColors(false)).Log()
	ctx := context.Background()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.InfoContext(ctx, testLog,
				`parm1`, `parm2`, `parm3`, `parm4`)
		}
	})
}

func BenchmarkJsonLogInfo(b *testing.B) {
	lg := NewLog(
		WithLevel(INFO),
		WithOutput(OutJson),
		WithStdOut(ioutil.Discard),
		WithFilePath(false),
		WithColors(false)).Log()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Info(testLog)
		}
	})
}

func BenchmarkJsonLogInfoFilePath(b *testing.B) {
	lg := NewLog(
		WithLevel(INFO),
		WithOutput(OutJson),
		WithStdOut(ioutil.Discard),
		WithFilePath(true),
		WithColors(false)).Log()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.Info(testLog)
		}
	})
}

func BenchmarkJsonInfoContext(b *testing.B) {
	lg := NewLog(
		WithLevel(INFO),
		WithOutput(OutJson),
		WithStdOut(ioutil.Discard),
		WithFilePath(false),
		WithColors(false)).Log()
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
	lg := NewLog(
		WithLevel(INFO),
		WithStdOut(ioutil.Discard),
		WithFilePath(false),
		WithOutput(OutJson),
		WithCtxExtractor(func(ctx context.Context) []interface{} {
			return []interface{}{ctx.Value(`ctx1`), ctx.Value(`ctx2`)}
		}),
		WithColors(false)).Log()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.InfoContext(ctx2, testLog)
		}
	})
}

func BenchmarkJsonInfoParams(b *testing.B) {
	lg := NewLog(
		WithLevel(INFO),
		WithOutput(OutJson),
		WithStdOut(ioutil.Discard),
		WithFilePath(false),
		WithColors(false)).Log()
	ctx := context.Background()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			lg.InfoContext(ctx, testLog, `parm1`, `parm2`, `parm3`, `parm4`)
		}
	})
}
