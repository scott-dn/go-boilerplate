package configs

func newDevelopmentConfig() *Config {
	return &Config{
		GoENV:    "development",
		HttpPort: 8080,
		PgDbURL:  "", // TODO: update here
	}
}
