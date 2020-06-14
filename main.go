package main

import (
	"log"

	"github.com/huntsman90/litserver/pkg/service"
)

func main() {
	s := service.NewLiteratureService()
	log.Println("Literature service is started.")
	log.Fatal(s.Server().ListenAndServe())
}
