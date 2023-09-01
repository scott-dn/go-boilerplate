package configs

func newProductionConfig() *Config {
	return &Config{
		GoENV:    "production",
		HttpPort: 8080,
		PgDbURL:  "", // TODO: update here
	}
}
