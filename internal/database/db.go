package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDataBase(dataSourceName string) (*Database, error) {
	db, err := sqlx.Connect("pgx", dataSourceName)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
