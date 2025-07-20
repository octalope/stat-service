package util

import (
	// "io"
	"os"
	"runtime"
	// "runtime/debug"
	// "strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var once sync.Once

var log zerolog.Logger

func Log() zerolog.Logger {
	once.Do(func() {
		// zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		// zerolog.TimeFieldFormat = time.RFC3339Nano

		// logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		// if err != nil {
		//     logLevel = int(zerolog.InfoLevel) // default to INFO
		// }

		// var output io.Writer = zerolog.ConsoleWriter{
		//     Out:        os.Stdout,
		//     TimeFormat: time.RFC3339,
		// }

		// if os.Getenv("APP_ENV") != "development" {
		// }

		// var gitRevision string

		// buildInfo, ok := debug.ReadBuildInfo()
		// if ok {
		//     for _, v := range buildInfo.Settings {
		//         if v.Key == "vcs.revision" {
		//             gitRevision = v.Value
		//             break
		//         }
		//     }
		// }

		log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			Level(zerolog.TraceLevel).
			With().
			Timestamp().
			Caller().
			Int("pid", os.Getpid()).
			Str("go_version", runtime.Version()).
			Logger()
	})

	return log
}
