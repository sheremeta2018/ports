package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tsetsik/ports/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")

	log.Printf("The server is running on port %s", port)

	router := mux.NewRouter()

	router.HandleFunc("/ports", handlers.GetPortsHelper).Methods("GET")
	router.HandleFunc("/ports", handlers.ImportPortsHelper).Methods("POST")

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
