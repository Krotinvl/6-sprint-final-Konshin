package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Loger  *log.Logger
	Server *http.Server
}

func StartServer(loger *log.Logger) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetHtml)
	mux.HandleFunc("/upload", handlers.HandlerForForm)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     loger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Server: srv,
		Loger:  loger,
	}
}
