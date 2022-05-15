package sqlc

import (
	"context"
	"database/sql"
)

var _ DBTX = (*wrappedDB)(nil)

func Wrap(db DBTX) DBTX {
	return &wrappedDB{db}
}

type wrappedDB struct {
	DBTX
}

func (w wrappedDB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.ExecContext(ctx, query, args...)
}

func (w wrappedDB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.QueryContext(ctx, query, args...)
}

func (w wrappedDB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if b, ok := BuilderFrom(ctx); ok {
		query, args = b.Build(query, args...)
	}

	return w.DBTX.QueryRowContext(ctx, query, args...)
}
