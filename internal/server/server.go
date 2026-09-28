package server

import (
	"net/http"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

func StartServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.GetHtml)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		return err
	}

	return nil
}
