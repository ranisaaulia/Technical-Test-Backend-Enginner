package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	username string = "postgres"
	password string = "admin"
	database string = "CRM_Test"
	host     string = "localhost"
	port     int    = 5432
)

var (
	dsn = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable",
		username, password, database, host, port)
)

func PostgreSQL() (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
