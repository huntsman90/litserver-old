package create

import (
	"log"
	"net/http"

	"github.com/huntsman90/litserver/pkg/gameroom"
)

// Create is the http Handle struct for creating gamerooms
type Create struct {
	GameRegister register
}

type register interface {
	Register(*gameroom.Gameroom) error
}

func (c Create) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gr := gameroom.New()
	err := c.GameRegister.Register(gr)
	if err != nil {
		log.Printf("game registry failed with error: %s", err.Error())
		return
	}
	log.Fprintln(w, "game room ID: ", gr.ID())
}
