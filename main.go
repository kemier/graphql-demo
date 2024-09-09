package main

import (
	"flag"
	"graphql-demo/client"
	"graphql-demo/server"
	"log"
	"os"
)

func main() {
	var command string
	flag.StringVar(&command, "command", "", "The command to run (server, client)")
	flag.Parse()

	switch command {
	case "server":
		if err := server.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	case "client":
		client.Query()
	default:
		log.Fatalf("Unknown command: %s", command)
		os.Exit(1)
	}
}
