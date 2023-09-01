package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/rs/zerolog/log"

	// Driver for migration.
	// This is needed to load the file source.
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func runMigration(instance *sql.DB) {
	driver, err := postgres.WithInstance(instance, &postgres.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("failed to create migration driver")
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Panic().Err(err).Msg("failed to get current working directory")
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s", filepath.Join(cwd, "migrations")),
		"postgres", driver)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create migration instance")
	}

	curVersion, _, err := migration.Version()
	if err != nil {
		log.Debug().Msg("current migration version is not found, set to 0")
		curVersion = 0
	}

	// NOTE:
	// It is not recommended to run the migration to the latest version
	// If you want to run the migration to the latest version,
	// you can use the following code:
	//
	// migration.Up()
	//
	toVersion := uint(1)
	if curVersion != toVersion {
		// Migrate looks at the currently active migration version,
		// then migrates either up or down to the specified version.
		if err := migration.Migrate(toVersion); err != nil {
			log.Panic().Err(err).Msg("failed to migrate database")
		}
	}
}
