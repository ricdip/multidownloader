package log

import (
	"multidownloader/src/flags"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	if flags.DebugLog {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else if flags.QuietLog {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:        os.Stderr,
			NoColor:    !flags.ColoredLog,
			TimeFormat: "01-02-06 15:04:05",
		}).
		With().
		Timestamp().
		Logger()
}
