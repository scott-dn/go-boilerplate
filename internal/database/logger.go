package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type dbLogger struct {
	slowThreshold time.Duration
}

func newDbLogger() *dbLogger {
	return &dbLogger{
		slowThreshold: 200 * time.Millisecond,
	}
}

func (l *dbLogger) LogMode( //nolint:ireturn
	gormlogger.LogLevel,
) gormlogger.Interface {
	return l
}

func (l *dbLogger) Info(_ context.Context, s string, args ...interface{}) {
	log.Info().Msgf(s, args...)
}

func (l *dbLogger) Warn(_ context.Context, s string, args ...interface{}) {
	log.Warn().Msgf(s, args...)
}

func (l *dbLogger) Error(_ context.Context, s string, args ...interface{}) {
	log.Error().Msgf(s, args...)
}

func (l *dbLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rowsAffected := fc()
	fields := map[string]interface{}{
		"duration":      fmt.Sprintf("%v ms", elapsed.Milliseconds()),
		"rows_affected": rowsAffected,
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error().Err(err).Fields(fields).Msgf("[GORM] query error: %s", sql)
		return
	}

	if l.slowThreshold != 0 && elapsed > l.slowThreshold {
		log.Error().Fields(fields).Msgf("[GORM] slow query: %s", sql)
		return
	}

	log.Debug().Fields(fields).Msgf("[GORM] query: %s", sql)
}
