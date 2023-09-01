package configs

func newProductionConfig() *Config {
	return &Config{
		GoENV:    "production",
		HTTPPort: 8080,
		PgDbURL:  "",         // TODO: update here
		CORS:     []string{}, // TODO: update here
	}
}
