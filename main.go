package main

import (
	"log"

	"github.com/cryptola-co/nail-server/cmd"
)

func main() {
	log.SetPrefix("nail ")
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
