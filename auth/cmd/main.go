package main

import (
	"database/sql"
	"log"

	"github.com/FadyGamilM/rhythmify/auth/api"
	"github.com/FadyGamilM/rhythmify/auth/business/auth"
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
	authService := auth.NewUserAuthService(userRepo)
	userService := user.NewUserService(userRepo)
	handlers := api.NewHandler(authService, userService)
	handlers.SetupEndpoints()
	server := api.Server(handlers)
	api.Run(server)
}
