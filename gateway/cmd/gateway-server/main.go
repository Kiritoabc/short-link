package main

import (
	"github.com/Kiritoabc/short-link/gateway/cmd/gateway-server/gateway"
	"log"
)

func main() {
	cmd := gateway.Command()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
