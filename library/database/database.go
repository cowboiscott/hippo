package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"github.com/cowboiscott/hippo/library/config"
)

type Database struct {
	db *sql.DB
}

func NewDB(config config.Config) (*Database, error) {
	db, err := sql.Open("sqlite3", config.Database.Path)
	if err != nil {
		return nil, err
	}

	return &Database{
		db,
	}, nil
}
