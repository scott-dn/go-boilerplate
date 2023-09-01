package configs

func newProductionConfig() *Config {
	return &Config{
		GoENV:     "production",
		HTTPPort:  8080,
		PgDbURL:   "postgres://service:password@localhost:5432/book?sslmode=disable", // TODO: update here
		JwtSecret: []byte("secret"),                                                  // TODO: update here
		CORS:      []string{"*"},                                                     // TODO: update here
	}
}
