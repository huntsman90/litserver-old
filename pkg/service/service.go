package service

import "net/http"

// NewLiteratureService is a constructor that returns literatureService
func NewLiteratureService() *LiteratureService {
	return &LiteratureService{server: &http.Server{Addr: ":8080"}}
}

type LiteratureService struct {
	server *http.Server
}

func (ls *LiteratureService) Server() *http.Server {
	return server
}
