package configs

func newLocalConfig() *Config {
	return &Config{
		GoENV:     "local",
		HTTPPort:  8080,
		PgDbURL:   "postgres://service:password@localhost:5432/book?sslmode=disable",
		JwtSecret: []byte("secret"),
		CORS:      []string{"*"},
	}
}
