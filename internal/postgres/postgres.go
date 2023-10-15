package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	*sql.DB
}

func Init(connStr string) (*DB, error) {
	db, dbErr := sql.Open("postgres", connStr)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	return &DB{db}, nil
}
