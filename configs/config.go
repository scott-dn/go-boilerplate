package configs

import (
	"os"
)

type Config struct {
	GoENV    string
	HttpPort uint
	PgDbURL  string
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
