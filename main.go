package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tsetsik/ports/internal/parser"
)

// Response struct for formin the actual json response
type Response struct {
	Error      bool        `json:"error"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
	statusCode int
}

// GetPortsHelper handler for getting all stored ports
func GetPortsHelper(w http.ResponseWriter, req *http.Request) {
	respond(http.StatusNotImplemented, nil, "", w)
}

// ImportPortsHelper handler for make the actual import
func ImportPortsHelper(w http.ResponseWriter, req *http.Request) {
	jsonPath := req.FormValue("file")

	if err := parser.ParseFileAsync(jsonPath); err != nil {
		respond(http.StatusBadRequest, nil, "", w)
	} else {
		respond(http.StatusAccepted, nil, "", w)
	}
}

func respond(status int, data interface{}, msg string, w http.ResponseWriter) {
	response := &Response{Error: false, Data: data, Message: msg, statusCode: status}

	if status >= 400 {
		response.Error = true
		response.Message = "Something went wrong"
	} else {
		response.Data = data
	}

	w.WriteHeader(response.statusCode)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err.Error())
	}

	port := os.Getenv("PORT")

	fmt.Printf("The server is running on port %s", port)

	router := mux.NewRouter()

	router.HandleFunc("/ports", GetPortsHelper).Methods("GET")
	router.HandleFunc("/ports", ImportPortsHelper).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
