package parser

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	client "github.com/tsetsik/ports-storage/pkg/client/v1"
)

// Parser interface
type Parser interface {
	ParseFileAsync(path string) error
}

type parser struct {
	client client.PortsStorageClient
}

// NewParser initiates new parser
func NewParser(clientURI string) (Parser, error) {
	client, err := client.NewClient(clientURI)
	if err != nil {
		return nil, err
	}

	return &parser{client: client}, nil
}

// ParseFileAsync that parse json file and return error
func (p *parser) ParseFileAsync(path string) error {
	input, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to read path %v with the following error: %v", path, err)
		return err
	}

	go p.processFile(input)

	return nil
}

func (p *parser) processFile(input *os.File) {
	dec := json.NewDecoder(input)

	t, _ := dec.Token()
	if delim, ok := t.(json.Delim); !ok || delim != '{' {
		log.Printf("Unknown starting token %v", delim)
		return
	}

	// while the object contains keys
	for dec.More() {
		portKey, err := dec.Token()
		if err != nil {
			log.Printf("Error in reading json token - %v with err: %v", portKey, err)
		}

		port := new(client.Port)
		err = dec.Decode(port)
		if err != nil {
			log.Printf("Error in decoding port: %v", err)
		}

		// Use code as ID
		i, err := strconv.Atoi(port.Code)
		if err != nil {
			log.Printf("Error in converting code field to id with the following error: %v", err)
			continue
		}
		port.ID = int32(i)

		// Store the parsed port in the db
		if _, err := p.client.UpsertPort(port); err != nil {
			log.Printf("Error in storing the port from the client: %v", err)
		}
	}
}
