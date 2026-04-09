//go:build cgo

package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func initDB() (*sql.DB, func(), error) {
	primaryURL := os.Getenv("TURSO_DATABASE_URL")
	authToken := os.Getenv("TURSO_AUTH_TOKEN")

	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create temp dir: %w", err)
	}

	dbPath := filepath.Join(dir, "cashflow.db")

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryURL,
		libsql.WithAuthToken(authToken),
	)
	if err != nil {
		os.RemoveAll(dir)
		return nil, nil, fmt.Errorf("failed to create connector: %w", err)
	}

	db := sql.OpenDB(connector)

	cleanup := func() {
		db.Close()
		connector.Close()
		os.RemoveAll(dir)
	}

	return db, cleanup, nil
}
