package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	var srvLogger log.Logger
	srv := server.CreateRouter(srvLogger)
	err := srv.Server.ListenAndServe()

	if err != nil {
		srv.Logger.Fatal(err)
	}
}
