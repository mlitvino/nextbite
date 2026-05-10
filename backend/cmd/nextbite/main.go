package main

import (
	"log"

	"github.com/mlitvino/nextbite/backend/internal/server"
)

func main() {
	srv := server.New()
	if err := srv.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
