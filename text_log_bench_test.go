package log_test

import (
	"context"
	"io/ioutil"
	"log"
	"testing"

	"github.com/google/uuid"
	lg "github.com/tryfix/log"
	traceable_context "github.com/tryfix/traceable-context"
)

var txtLg = lg.NewLog(lg.WithLevel(lg.INFO), lg.WithStdOut(ioutil.Discard), lg.WithFilePath(true), lg.WithColors(true))
var txtLogger = txtLg.Log()
var native = log.New(ioutil.Discard, `test`, log.Lmicroseconds)

var txtTestCtx = traceable_context.WithUUID(uuid.New())

func BenchmarkTextNative(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			native.Println(`dd`)
		}
	})
}

func BenchmarkTextInfo(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			txtLogger.Info(`dd`)
		}
	})
}

func BenchmarkTextInfoContext(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			txtLogger.InfoContext(txtTestCtx, `dd`)
		}
	})
}

func BenchmarkTextInfoParams(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			other := txtLogger.NewLog(lg.Prefixed(`ss`))
			go other.InfoContext(context.Background(), `dd`, 1, 2, 3)
			xx := other.NewLog(lg.Prefixed(`ss`))
			go xx.InfoContext(context.Background(), `dd`, 1, 2, 3)
		}
	})
}
