package app

import (
	"github.com/rs/zerolog/log"
	"github.com/scott-dn/go-boilerplate/configs"
	"github.com/scott-dn/go-boilerplate/internal/database"
	"github.com/scott-dn/go-boilerplate/internal/pkg/logger"
	"github.com/scott-dn/go-boilerplate/internal/service"
	"gorm.io/gorm"
)

type App struct {
	database *gorm.DB

	Config  *configs.Config
	Service *service.Service
}

func (app *App) Close() {
	log.Info().Msg("closing database connection")
	db, err := app.database.DB()
	if err != nil {
		log.Panic().Err(err).Msg("failed to get database connection")
	}
	err = db.Close()
	if err != nil {
		log.Panic().Err(err).Msg("failed to close database connection")
	}
}

func (app *App) HealthCheck() error {
	var result int64
	if err := app.database.Raw("SELECT 1").Scan(&result).Error; err != nil || result != 1 {
		log.Error().Err(err).Msg("database health check failed")
		return err
	}
	/*
	 * if err := app.Cache.Ping(); err != nil {
	 *   log.Error().Err(err).Msg("cache health check failed")
	 *   return err
	 * }
	 */
	return nil
}

func Init() *App {
	config := configs.NewConfig()
	logger.InitGlobal(config)

	log.Info().
		Any("config", config).
		Msg("config initialized")

	database := database.Init(config)
	log.Info().Msg("init database successfully")

	return &App{
		database: database,
		Config:   config,
		Service:  service.NewService(database),
	}
}
