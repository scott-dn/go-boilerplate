package configs

func newUATConfig() *Config {
	return &Config{
		GoENV:    "uat",
		HTTPPort: 8080,
		PgDbURL:  "",         // TODO: update here
		CORS:     []string{}, // TODO: update here
	}
}
