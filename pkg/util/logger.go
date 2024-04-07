package util

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func SetupLogger(debug *bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	levelFromConfig := GetConfigString("logging.level")
	if level, err := zerolog.ParseLevel(levelFromConfig); err != nil {
		log.Error().Err(err).Msgf("Error parsing log level: %s", err)
		panic(err)
	} else {
		zerolog.SetGlobalLevel(level)
	}

	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func GetLogger() *zerolog.Logger {
	return &log.Logger
}

/*
panic (zerolog.PanicLevel, 5)
fatal (zerolog.FatalLevel, 4)
error (zerolog.ErrorLevel, 3)
warn (zerolog.WarnLevel, 2)
info (zerolog.InfoLevel, 1)
debug (zerolog.DebugLevel, 0)
trace (zerolog.TraceLevel, -1)

Global level that is set will log all messages with a level greater or equal to the global level. i.e. if the global level is set to info, all messages with a level of info, warn, error, fatal, and panic will be logged. If the global level is set to debug, all messages with a level of debug, info, warn, error, fatal, and panic will be logged.
*/
