package helloworld

import (
	"log"
	"net/http"

	"github.com/huntsman90/litserver/pkg/gameroom"
)

// HelloWorld is the handler struct for /helloworld route.
type HelloWorld struct {
	GameRegister register
}

type register interface {
	Register(*gameroom.Gameroom) error
}

func (hw HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gr := gameroom.New()
	err := hw.GameRegister.Register(gr)
	if err != nil {
		log.Printf("game registry failed with error: %s", err.Error())
		return
	}
}
