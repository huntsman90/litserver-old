package play

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/huntsman90/litserver/pkg/gameroom"
)

// Play is the HTTP handler struct fo the /play endpoint.
type Play struct {
	GameRegister register
}

type register interface {
	Room(string) (*gameroom.Gameroom, error)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(_ *http.Request) bool {
		return true
	},
}

func (p Play) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}

	gameID := r.FormValue("id")
	if gameID == "" {
		log.Println("missing game ID in request")
		return
	}

	room, err := p.GameRegister.Room(gameID)
	if err != nil {
		log.Println(err.Error())
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Client connected.")

	readFrom, sendTo := room.Channels()
	listenToRoom(readFrom, conn)

	reader(sendTo, conn)
}

func reader(sender chan string, conn *websocket.Conn) {
	for {
		// read in a message
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		sender <- string(p)

		log.Println("Message received: ", string(p))
	}
}

func listenToRoom(receiver chan string, conn *websocket.Conn) {
	go func() {
		for {
			select {
			case msg := <-receiver:
				log.Println("Received message from room: ", msg)
				if err := conn.WriteMessage(websocket.BinaryMessage, []byte(msg)); err != nil {
					log.Println(err)
					return
				}
			}
		}

	}()
}
