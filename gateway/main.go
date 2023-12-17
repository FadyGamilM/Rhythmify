package main

import (
	"log"

	"github.com/FadyGamilM/rhythmify/gateway/api"
	mongogridfs "github.com/FadyGamilM/rhythmify/gateway/mongo-gridfs"
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
	gridFs, err := mongogridfs.Connect()
	if err != nil {
		log.Fatalf("couldn't connect to mongodb ... ")
	}
	router := api.NewHandler(gridFs)
	router.SetupEndpoints()
	server := api.Server(router)
	api.Run(server)
}
