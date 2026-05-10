package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mlitvino/nextbite/backend/internal/api"
)

func main() {
	r := gin.Default()
	api.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
