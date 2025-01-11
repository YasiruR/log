package main

import (
	"context"
	"github.com/YasiruR/log"
	"github.com/google/uuid"
	traceableContext "github.com/tryfix/traceable-context"
)

func main() {
	sl := log.NewLog(
		log.WithCtxTraceExtractor(func(ctx context.Context) string {
			if trace := traceableContext.FromContext(ctx); trace != uuid.Nil {
				return trace.String()
			}

			return ""
		}),
	).Log()
	//l := log.NewLog(log.Prefixed(`prefix1`)).Log()
	//pl := log.NewLog().PrefixedLog()

	sl.Info("logged a simple message with params", "param1", "param2")
	//l.Info("logger initialized with static prefix", "param1", "param2")
	//pl.Info(`prefix2`, `prefixed-logger message with params`, "param1", "param2")

	//l.Info("this is a sample message")
	//pl.Info(`sample prefix`, `a message`)
}
