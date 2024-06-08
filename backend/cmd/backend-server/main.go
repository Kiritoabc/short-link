package main

import (
	"github.com/Kiritoabc/short-link/backend/cmd/backend-server/app"
	"log"
)

func main() {
	cmd := app.Command()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
