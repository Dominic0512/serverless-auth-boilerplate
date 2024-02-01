package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	app, err := InitializeApp()
	if err != nil {
		log.Fatal(fmt.Printf("failed to create route: %s\n", err))
		os.Exit(2)
	}

	app.Start()
}
