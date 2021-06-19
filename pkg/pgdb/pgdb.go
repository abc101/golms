package pgdb

import (
	"database/sql"
	"fmt"

	"github.com/abc101/golms/pkg/config"
)

// Connect will returns sql.DB when , otherwise error
func Connect(pginfo config.Postgres) (dbconn *sql.DB, err error) {

	// Postgres connection statement
	pgconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", pginfo.Host, pginfo.Port, pginfo.User, pginfo.Password, pginfo.Dbname, pginfo.SSLmode)

	db, err := sql.Open("postgres", pgconn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
