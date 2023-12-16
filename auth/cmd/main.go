package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/FadyGamilM/rhythmify/auth/api"
	"github.com/FadyGamilM/rhythmify/auth/business/auth"
	"github.com/FadyGamilM/rhythmify/auth/business/user"
	"github.com/FadyGamilM/rhythmify/auth/db/postgres"
	"github.com/FadyGamilM/rhythmify/auth/initializer"
	"github.com/FadyGamilM/rhythmify/auth/repository"
)

var err error
var db *sql.DB

var (
	postgres_host    = os.Getenv("POSTGRES_HOST")
	postgres_user    = os.Getenv("POSTGRES_USER")
	postgres_pass    = os.Getenv("POSTGRES_PASSWORD")
	postgres_db      = os.Getenv("POSTGRES_DB")
	postgres_sslmode = os.Getenv("POSTGRES_SSLMODE")
)

// runs before main
func init() {

	postgres_host = os.Getenv("POSTGRES_HOST")
	postgres_user = os.Getenv("POSTGRES_USER")
	postgres_pass = os.Getenv("POSTGRES_PASSWORD")
	postgres_db = os.Getenv("POSTGRES_DB")
	postgres_sslmode = os.Getenv("POSTGRES_SSLMODE")

	err := initializer.LoadEnvVars()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// connect to database before anything
	db, err = postgres.SetupConnection()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	postgresDB := postgres.NewPG(db)
	userRepo := repository.NewPgUserRepo(postgresDB)
	authService := auth.NewUserAuthService(userRepo)
	userService := user.NewUserService(userRepo)
	handlers := api.NewHandler(authService, userService)
	handlers.SetupEndpoints()
	server := api.Server(handlers)
	api.Run(server)
}
