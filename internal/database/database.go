package database

import (
	"database/sql"
	"fmt"

	"lastpath/internal/entity"
)

type DB struct {
	db *sql.DB
}

func NewDB() *DB {
	var db DB
	return &db
}

func (db *DB) Create(path *entity.Path) error {
	fmt.Println(path)
	return nil
}
