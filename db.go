package sqlc

import (
	"context"
	"database/sql"

	"github.com/yiplee/nap"
)

var _ DBTX = (*DB)(nil)

type DB struct {
	*nap.DB
}

func Connect(driverName, master string, slaves ...string) (*DB, error) {
	db, err := nap.Open(driverName, master, slaves...)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// PrepareContext creates a prepared statement for later queries or executions.
//
// Read or update can not be determined by the query string currently, use master database.
func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.Master().PrepareContext(ctx, query)
}
