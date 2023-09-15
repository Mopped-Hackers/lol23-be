package main

import (
	"log"
	"lol23-be/endpoints"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := gin.Default()
	routes(r)
	r.Run()
}

func routes(r *gin.Engine) {
	r.GET("/health", endpoints.Health)
}
