package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mlitvino/nextbite/backend/internal/config"
	"github.com/mlitvino/nextbite/backend/internal/server"
)

func main() {
    if err := godotenv.Load("../.env"); err != nil {
        log.Fatal(err)
    }

    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    srv := server.New()
    if err := srv.Run(":" + cfg.Port); err != nil {
        log.Fatal(err)
    }
}
