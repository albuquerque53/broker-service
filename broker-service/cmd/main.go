package main

import (
	"broker/internal/application/route"
	"fmt"
	"log"
	"net/http"
)

const webPort = "80"

func main() {
	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: route.Routes(),
	}

	err := srv.ListenAndServe()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
