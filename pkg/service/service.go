package service

import (
	"net/http"

	"github.com/huntsman90/litserver/pkg/play"

	"github.com/huntsman90/litserver/pkg/gameroom/create"

	"github.com/huntsman90/litserver/pkg/gameroom/registry"
)

// NewLiteratureService is a constructor that returns literatureService
func NewLiteratureService() *LiteratureService {
	rg := registry.New()
	mux := http.NewServeMux()
	mux.Handle("/create", create.Create{GameRegister: rg})
	mux.Handle("/play", play.Play{GameRegister: rg})
	// mux.Handle("/helloworld", helloworld.HelloWorld{GameRegister: rg})
	// mux.Handle("/hellosocket", hellosocket.HelloSocket{GameRegister: rg})

	return &LiteratureService{server: &http.Server{Addr: ":8080", Handler: mux}}
}

// LiteratureService is the struct that contains the outward facing APIs to conduct a game of literature
type LiteratureService struct {
	server *http.Server
}

// Server returns the HTTP server that handles requests to the service.
func (ls *LiteratureService) Server() *http.Server {
	return ls.server
}
