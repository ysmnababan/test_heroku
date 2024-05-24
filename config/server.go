package config

import (
	"fmt"
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
	if port == "" {
		port = "8080"
	}
	server := http.Server{
		Addr:    "localhost" + fmt.Sprintf(":%s", port),
		Handler: router,
	}

	return router, &server
}
