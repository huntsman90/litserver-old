package main

import (
	"log"
)

func main() {
	service := NewLiteratureService()
	log.Fatal(service.Server.ListenAndServe())
}
