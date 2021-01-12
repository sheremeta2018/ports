package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// Port is
type Port struct {
	Name     string `json:"name"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Province string `json:"provin"`
}

// ParseFileAsync that parse json file and return error
func ParseFileAsync(path string) error {
	input, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to read path %v with the following error: %v", path, err)
		return err
	}

	go processFile(input)

	return nil
}

func processFile(input *os.File) {
	dec := json.NewDecoder(input)

	for {
		var v map[string]*Port
		if err := dec.Decode(&v); errors.Is(err, io.EOF) {
			break // done decoding file
		} else if err != nil {
			log.Fatal(err)
		}

		for code, data := range v {
			fmt.Printf("\n Processing code %v with data %v", code, data)
			// We should use the ports-storage client and store each port
		}
	}
}
