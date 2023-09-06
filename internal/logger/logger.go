package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/scott-dn/go-boilerplate/configs"
)

func InitGlobal(config *configs.Config) {
	switch config.GoENV {
	case "local":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	default:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.TimestampFieldName = "t"
		zerolog.LevelFieldName = "l"
		zerolog.MessageFieldName = "m"
		zerolog.ErrorFieldName = "e" //nolint:reassign
	}
}
