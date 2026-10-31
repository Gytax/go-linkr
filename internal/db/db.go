package db

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type DB struct {
	path string
	db   *sqlx.DB
}

func Open(path string) (*DB, error) {
	if path == "" {
		return nil, fmt.Errorf("database path is empty")
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("create database directory: %w", err)
	}

	sqlDB, err := sqlx.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	for _, query := range []string{
		"PRAGMA journal_mode=WAL;",
		"PRAGMA synchronous=NORMAL;",
		"PRAGMA busy_timeout = 5000;",
	} {
		if _, err := sqlDB.Exec(query); err != nil {
			sqlDB.Close()
			return nil, fmt.Errorf("configure sqlite pragmas (%s): %w", query, err)
		}
	}

	db := &DB{path: path, db: sqlDB}
	if err := db.initializeSchema(); err != nil {
		sqlDB.Close()
		return nil, err
	}

	return db, nil
}

func (d *DB) Close() error {
	if d == nil || d.db == nil {
		return nil
	}
	return d.db.Close()
}

func (d *DB) Path() string {
	if d == nil {
		return ""
	}
	return d.path
}
