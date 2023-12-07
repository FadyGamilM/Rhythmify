package main

import (
	"database/sql"
	"log"

	"github.com/FadyGamilM/rhythmify/auth/api"
	"github.com/FadyGamilM/rhythmify/auth/business/user"
	"github.com/FadyGamilM/rhythmify/auth/db/postgres"
	"github.com/FadyGamilM/rhythmify/auth/initializer"
	"github.com/FadyGamilM/rhythmify/auth/repository"
)

var db *sql.DB

// runs before main
func init() {
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
	userService := user.NewUserAuthService(userRepo)
	handlers := api.NewHandler(userService)
	handlers.SetupEndpoints()
	server := api.Server(handlers)
	api.Run(server)
}
