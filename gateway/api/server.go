package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Server(h *Handler) *http.Server {
	port := os.Getenv("GATEWAY_SERVER")
	log.Printf("the port is : {%v}", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", port),
		Handler: h.router,
	}
	return srv
}

func Run(srv *http.Server) {
	log.Println("gateway server is up and running on : ", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server is down : %v", err)
	}
}
