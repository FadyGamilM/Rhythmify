package main

import (
	"log"

	"github.com/FadyGamilM/rhythmify/auth/api"
	"github.com/FadyGamilM/rhythmify/auth/initializer"
)

// runs before main
func init() {
	err := initializer.LoadEnvVars()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func main() {
	handlers := api.NewHandler()
	handlers.SetupEndpoints()
	server := api.Server(handlers)
	api.Run(server)
}
