package main

import (
	"log"

	"github.com/FadyGamilM/rhythmify/gateway/api"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	router := api.NewHandler()
	router.SetupEndpoints()
	server := api.Server(router)
	api.Run(server)
}
