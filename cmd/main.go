package main

import (
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	if err := server.StartServer(); err != nil {
		fmt.Printf("Error starting server: %s", err)
	}

}
