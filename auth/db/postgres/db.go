package postgres

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

// this interface type is used as a dependency in all repos so any repo can receive a transaction "tx" or the database pool itself "db" and execute the transaction via anyone of them because both implements DBTX
type DBTX interface {
	// i named the params here because i need the users to know what they should pass to this func
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)

	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)

	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)

	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// a wrapper used by all repos to access the database instance
type PG struct {
	// this type compose one property which is of type DBTX
	// this type allows my struct to accepts either tranxation *sql.TX or single query *sql.DB
	DB DBTX
}

// factory method
func NewPG(db DBTX) *PG {
	return &PG{
		DB: db,
	}
}
