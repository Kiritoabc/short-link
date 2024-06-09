package main

import (
	"github.com/Kiritoabc/short-link/backend/cmd/backend-server/app"
	_ "github.com/Kiritoabc/short-link/docs"
	"log"
)

func main() {
	cmd := app.Command()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
