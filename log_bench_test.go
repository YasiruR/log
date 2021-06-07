package log

import (
	"bytes"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	traceable_context "github.com/tryfix/traceable-context"
)

var byt = bytes.NewBuffer(make([]byte, 100))
var lg = NewLog(WithLevel(INFO), WithStdOut(byt), WithFilePath(true), WithColors(true))
var pxLg = lg.Log()
var native = log.New(byt, `test`, log.Lmicroseconds)

var testCtx = traceable_context.WithUUID(uuid.New())

func BenchmarkNative(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			native.Println(`dd`)
		}
	})
}

func BenchmarkInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pxLg.Info(`dd`)
		}
	})
}

func BenchmarkInfoContext(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pxLg.InfoContext(testCtx, `dd`)
		}
	})
}

func BenchmarkInfoParams(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			other := pxLg.NewLog(Prefixed(`ss`))
			go other.InfoContext(context.Background(), `dd`, 1, 2, 3)
			xx := other.NewLog(Prefixed(`ss`))
			go xx.InfoContext(context.Background(), `dd`, 1, 2, 3)
		}
	})
}
