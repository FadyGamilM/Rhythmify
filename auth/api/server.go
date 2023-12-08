package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Server(h *Handler) *http.Server {
	port := os.Getenv("SERVER_PORT")
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: h.Router,
	}
	return srv
}

func Run(srv *http.Server) error {
	log.Println("starting auth microservice .. ")
	if err := srv.ListenAndServe(); err != nil {
		return errors.New("failed to start the server")
	}
	return nil
}
