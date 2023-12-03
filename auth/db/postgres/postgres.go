package postgres

import (
	"database/sql"
	"log"
	"net/url"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

var (
	postgres_host    = os.Getenv("POSTGRES_HOST")
	postgres_user    = os.Getenv("POSTGRES_USER")
	postgres_pass    = os.Getenv("POSTGRES_PASSWORD")
	postgres_db      = os.Getenv("POSTGRES_DB")
	postgres_sslmode = os.Getenv("POSTGRES_SSLMODE")
)

func SetupConnection() (*sql.DB, error) {
	// get the dsn
	dsn := getDSN(postgres_host, postgres_user, postgres_pass, postgres_db)

	q := dsn.Query()
	q.Add("sslmode", postgres_sslmode)

	// get a url encoded for postgres db
	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		log.Printf("error trying to open a postgres connection ➜ %v \n", err)
		return nil, err
	}

	log.Println("connected to postgres instance successfully .. ")

	return db, nil
}

func getDSN(host, user, password, dbname string) url.URL {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   host,
		User:   url.UserPassword(user, password),
		Path:   dbname,
	}

	return dsn
}
