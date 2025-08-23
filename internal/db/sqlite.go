package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func NewSQLite3() (*sql.DB, error) {
	path := "./demo.db"
	zap.S().Debugf("Starting connection with SQLite3, with db %s", path)

	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, fmt.Errorf("error creating file connection")
	}

	// Simple ping
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error while pinging (should never fail since its a file)")
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON`); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("enable foreign_keys: %w", err)
	}

	return db, nil
}
