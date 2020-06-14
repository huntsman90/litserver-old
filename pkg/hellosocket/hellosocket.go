package hellosocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// HelloSocket struct for the /hellosocket route. Allows us to pass the registry in.
type HelloSocket struct {
	GameRegister register
}

type register interface{}

// ServeHTTP is the handler for /hellosocket route. Satisfies the http.Handler interface.
func (HelloSocket) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		log.Println("/hellosocket websocket upgrade failed.")
		return
	}

	log.Println("Client connected.")
	reader(conn)

}

func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out that message for clarity
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
