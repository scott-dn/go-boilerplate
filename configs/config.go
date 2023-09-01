package configs

import (
	"os"
)

type Config struct {
	GoENV    string
	HTTPPort uint
	PgDbURL  string
	CORS     []string // whitelist
}

func NewConfig() *Config {
	switch os.Getenv("GO_ENV") {
	case "production":
		return newProductionConfig()
	case "uat":
		return newUATConfig()
	case "development":
		return newDevelopmentConfig()
	case "local":
		return newLocalConfig()
	default:
		return newLocalConfig()
	}
}
