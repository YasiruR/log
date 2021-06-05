package main

import (
	"github.com/tryfix/log"
)

func main() {
	// // usage of log
	// log.Info(`message`, `param1`, `param2`)
	// log.Trace(`message`)
	// log.Error(`message`)
	// log.Error(log.WithPrefix(`prefix`, `message`), `param1`, `param2`)

	// // log with a traceable context
	// tCtx := traceable_context.WithUUID(uuid.New())
	// ctx, _ := context.WithCancel(tCtx)
	// logger := log.Constructor.Log(log.WithColors(true), log.WithLevel(log.TRACE), log.WithFilePath(false), log.Prefixed(`level-1`))
	// logger.ErrorContext(ctx, `message`, `param1`, `param2`)
	// logger.ErrorContext(ctx, `message`)
	// logger.ErrorContext(ctx, `message`)
	// logger.ErrorContext(ctx, log.WithPrefix(`prefix`, `message`))
	// logger.WarnContext(ctx, log.WithPrefix(`prefix`, `message`), `param1`, `param2`)

	// // prefixed log
	// prefixedLogger := log.Constructor.PrefixedLog(log.WithLevel(log.ERROR), log.WithFilePath(true))
	// prefixedLogger.Info(`module.sub-module`, `message`)
	// prefixedLogger.Trace(`module.sub-module`, `message`)
	// prefixedLogger.Error(`module.sub-module`, `message`)
	// prefixedLogger.Error(`module.sub-module`, `message`, `param1`, `param2`)

	// // custom logger
	// //customLogger := customLog.NewLogger()
	// //customLogger.Info(`info`)
	// //customLogger.Trace(`trace`)
	// //
	// //// create a logger instance derived from logger
	// //nestedLogger := logger.NewLog(log.WithLevel(log.TRACE), log.Prefixed(`level-2`))
	// //nestedLogger.Error(`error happened`, 22)

	// // enable context reading
	// // keys
	// type keyOne string
	// type keyTwo string

	// const k1 keyOne = "key1"
	// const k2 keyTwo = "key2"

	// // get details from context
	// lCtx := context.Background()
	// lCtx = context.WithValue(lCtx, k1, "context_val_1")
	// lCtx = context.WithValue(lCtx, k2, "context_val_2")

	// ctxLogger := log.Constructor.Log(log.WithColors(true),
	// 	log.WithLevel(log.TRACE),
	// 	log.WithFilePath(false),
	// 	log.Prefixed(`context_logger`),
	// 	log.WithCtxExtractor(func(ctx context.Context) []interface{} {
	// 		return []interface{}{
	// 			fmt.Sprintf("%s: %+v", k1, ctx.Value(k1)),
	// 		}
	// 	}),
	// )

	// ctxLogger.ErrorContext(lCtx, `message`, `param1`, `param2`)
	// ctxLogger.ErrorContext(lCtx, `message`)
	// ctxLogger.ErrorContext(lCtx, `message`)
	// ctxLogger.ErrorContext(lCtx, log.WithPrefix(`prefix`, `message`))
	// ctxLogger.WarnContext(lCtx, log.WithPrefix(`prefix`, `message`), `param1`, `param2`)

	// // child logger with additional context extraction functionality
	// ctxChildLogger := ctxLogger.NewLog(log.Prefixed(`context_child_logger`),
	// 	log.WithCtxExtractor(func(ctx context.Context) []interface{} {
	// 		return []interface{}{
	// 			fmt.Sprintf("%s: %+v", k2, ctx.Value(k2)),
	// 		}
	// 	}),
	// )

	// ctxChildLogger.ErrorContext(lCtx, `message`, `param1`, `param2`)
	// ctxChildLogger.ErrorContext(lCtx, `message`)
	// ctxChildLogger.ErrorContext(lCtx, `message`)
	// ctxChildLogger.ErrorContext(lCtx, log.WithPrefix(`prefix`, `message`))
	// ctxChildLogger.WarnContext(lCtx, log.WithPrefix(`prefix`, `message`), `param1`, `param2`)

	// json logger
	jsonLogConstructor := log.NewLog(log.WithOutput(log.OutJson), log.WithFilePath(true))
	jsonLogger := jsonLogConstructor.Log()
	jsonLogger.Error("error message")
}
