package configs

func newDevelopmentConfig() *Config {
	return &Config{
		GoENV:     "development",
		HTTPPort:  8080,
		PgDbURL:   "postgres://service:password@localhost:5432/book?sslmode=disable", // TODO: update here
		JwtSecret: []byte("secret"),                                                  // TODO: update here
		CORS:      []string{"*"},                                                     // TODO: update here
	}
}
