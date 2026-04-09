//go:build !cgo

package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func initDB() (*sql.DB, func(), error) {
	db, err := sql.Open("sqlite", "./cashflow.db")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database: %w", err)
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}
