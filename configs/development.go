package configs

func newDevelopmentConfig() *Config {
	return &Config{
		GoENV:    "development",
		HTTPPort: 8080,
		PgDbURL:  "",         // TODO: update here
		CORS:     []string{}, // TODO: update here
	}
}
