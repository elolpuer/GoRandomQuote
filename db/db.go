package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db sql.DB

//Connect ...
func connect() error {
	var err error

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ", cfg.PgHost, cfg.PgPort, cfg.PgUser, cfg.PgPass, cfg.PgDB))
	if err != nil {
		return err
	}

	return nil
}
