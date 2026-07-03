package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

func initDB() (*sql.DB, func(), error) {
	tursoURL := os.Getenv("TURSO_DATABASE_URL")
	tursoToken := os.Getenv("TURSO_AUTH_TOKEN")

	if tursoURL != "" {
		dsn := tursoURL
		if tursoToken != "" {
			dsn = dsn + "?authToken=" + tursoToken
		}

		db, err := sql.Open("libsql", dsn)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to open Turso database: %w", err)
		}

		cleanup := func() {
			db.Close()
		}

		return db, cleanup, nil
	}

	db, err := sql.Open("sqlite", "./cashflow.db")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database: %w", err)
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}
