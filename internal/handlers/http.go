package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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

	//TODO: better to have config package from which we can access the variables
	clientURI := os.Getenv("PORTS_CLIENT_URI")

	p, err := parser.NewParser(clientURI)
	if err != nil {
		log.Printf("Error in initiating new parser with the following error: %v", err)
		respond(http.StatusBadRequest, nil, "Unable to parse the specified file", w)
		return
	}

	if err := p.ParseFileAsync(jsonPath); err != nil {
		log.Printf("Error in parsing %v with the following error: %v", jsonPath, err)
		respond(http.StatusBadRequest, nil, "Unable to parse the specified file", w)
	} else {
		respond(http.StatusAccepted, nil, "", w)
	}
}

func respond(status int, data interface{}, msg string, w http.ResponseWriter) {
	response := &Response{Error: false, Data: data, Message: msg, statusCode: status}

	if status >= 400 {
		response.Error = true
		response.Message = msg
	} else {
		response.Data = data
	}

	w.WriteHeader(response.statusCode)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
