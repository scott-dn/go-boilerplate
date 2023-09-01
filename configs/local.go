package configs

import "os"

func newLocalConfig() *Config {
	PgDbURL := "postgres://service:password@localhost:5432/service?sslmode=disable"
	overridePgDbURL := os.Getenv("PG_DB_URL")
	if overridePgDbURL != "" {
		PgDbURL = overridePgDbURL
	}
	return &Config{
		GoENV:    "local",
		HttpPort: 8080,
		PgDbURL:  PgDbURL,
	}
}
