package main

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/scott-dn/go-boilerplate/configs"
	"github.com/scott-dn/go-boilerplate/internal/database/entities"
	"github.com/scott-dn/go-boilerplate/internal/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Querier interface{}

func main() {
	config := configs.NewConfig()
	logger.InitGlobal(config)
	log.Info().
		Any("config", config).
		Msg("config initialized")

	gorm, err := gorm.Open(postgres.Open(config.PgDbURL), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("failed to connect database")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Panic().Err(err).Msg("failed to get current working directory")
	}

	generator := gen.NewGenerator(gen.Config{
		OutPath:           filepath.Join(cwd, "internal/database/query"),
		OutFile:           "gen.go",
		ModelPkgPath:      "query",
		WithUnitTest:      false,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  false,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	generator.UseDB(gorm)

	generator.ApplyBasic(
		entities.Book{},
	)

	generator.ApplyInterface(func(Querier) {},
		entities.Book{},
	)

	generator.Execute()
}
