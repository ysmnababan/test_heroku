package config

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func InitServer() (*httprouter.Router, *http.Server) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error while opening the .env file: %v", err)
	}

	router := httprouter.New()
	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return router, &server
}
