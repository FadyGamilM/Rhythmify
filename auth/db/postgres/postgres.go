package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

// var (
// 	postgres_host    = os.Getenv("POSTGRES_HOST")
// 	postgres_user    = os.Getenv("POSTGRES_USER")
// 	postgres_pass    = os.Getenv("POSTGRES_PASSWORD")
// 	postgres_db      = os.Getenv("POSTGRES_DB")
// 	postgres_sslmode = os.Getenv("POSTGRES_SSLMODE")
// )

func SetupConnection() (*sql.DB, error) {
	postgres_host := os.Getenv("POSTGRES_HOST")
	postgres_user := os.Getenv("POSTGRES_USER")
	postgres_pass := os.Getenv("POSTGRES_PASSWORD")
	postgres_db := os.Getenv("POSTGRES_DB")
	postgres_sslmode := os.Getenv("POSTGRES_SSLMODE")
	log.Println("======================== SETUP POSTGRES CONNECTION ========================")
	log.Printf("postgres host = {%v}\n", postgres_host)
	log.Printf("postgres user = {%v}\n", postgres_user)
	log.Printf("postgres pass = {%v}\n", postgres_pass)
	log.Printf("postgres dbname = {%v}\n", postgres_db)
	log.Printf("postgres sslmode = {%v}\n", postgres_sslmode)
	log.Println("===========================================================================")

	// get the dsn
	dsn := getDSN(postgres_host, postgres_user, postgres_pass, postgres_db)

	q := dsn.Query()
	q.Add("sslmode", postgres_sslmode)

	// get a url encoded for postgres db
	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		log.Println("[error in SetupConnection]")
		return nil, fmt.Errorf("error trying to open a postgres connection ➜ %v", err)
	}
	log.Println("connected to postgres instance successfully .. ")

	if err = db.Ping(); err != nil {
		log.Println("[error in SetupConnection]")
		return nil, fmt.Errorf("error trying to ping to postgres connection ➜ %v", err)
	}
	log.Println("pong .. ")

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
