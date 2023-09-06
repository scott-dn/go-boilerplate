package database

import (
	"github.com/rs/zerolog/log"
	"github.com/scott-dn/go-boilerplate/configs"
	"github.com/scott-dn/go-boilerplate/internal/database/query"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(config *configs.Config) *gorm.DB {
	database, err := gorm.Open(postgres.Open(config.PgDbURL), &gorm.Config{
		Logger:                 newDbLogger(),
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect database")
	}

	if sqlInstance, err := database.DB(); err != nil {
		log.Panic().Err(err).Msg("failed to get sql instance")
	} else {
		runMigration(sqlInstance)
	}

	query.SetDefault(database)

	return database
}
