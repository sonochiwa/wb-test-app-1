package main

import (
	"log"
	"net/http"

	"github.com/sonochiwa/wb-test-app/internal/handler"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler.InitRoutes(),
	}

	log.Println("Server running on http://0.0.0.0:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
