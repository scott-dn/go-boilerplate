package configs

func newUATConfig() *Config {
	return &Config{
		GoENV:     "uat",
		HTTPPort:  8080,
		PgDbURL:   "postgres://service:password@localhost:5432/book?sslmode=disable", // TODO: update here
		JwtSecret: []byte("secret"),                                                  // TODO: update here
		CORS:      []string{"*"},                                                     // TODO: update here
	}
}
