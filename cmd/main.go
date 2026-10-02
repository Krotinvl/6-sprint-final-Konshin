package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	var logger *log.Logger
	srv := server.StartServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		log.Fatalf("Server starting error: %s", err)
	}
}
