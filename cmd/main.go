package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ulule/helloapp"
)

func main() {
	// use PORT environment variable, or default to 8080
	rawPort := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		rawPort = fromEnv
	}

	port, err := strconv.Atoi(rawPort)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Server listening on port %d", port)

	err = helloapp.Run(port)
	if err != nil {
		log.Fatal(err)
	}
}
