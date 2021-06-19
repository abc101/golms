package config

// Config is a block of all specific configuration
type Config struct {
	Postgres Postgres
}

// Postgres connection
type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	SSLmode  string
}
