package configs

func newUATConfig() *Config {
	return &Config{
		GoENV:    "uat",
		HttpPort: 8080,
		PgDbURL:  "", // TODO: update here
	}
}

