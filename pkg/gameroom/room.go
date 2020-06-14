package gameroom

import (
	"log"
	"math/rand"
	"time"
	"unsafe"
)

// Gameroom object coordinates the communication between the players of the game.
type Gameroom struct {
	id       string
	senders  []chan string
	receiver chan string
}

// Channels method returns two channels:
//  send -> channel on which gameroom sends messages to the client
//  receive -> channel on which gameroom receives messages from the client
func (gr *Gameroom) Channels() (send chan string, receive chan string) {
	ch := make(chan string)
	gr.senders = append(gr.senders, ch)
	return ch, gr.receiver
}

// ID returns the identifier for the gameroom
func (gr *Gameroom) ID() string {
	return gr.id
}

func (gr *Gameroom) start() {
	go func() {
		for {
			select {
			case m := <-gr.receiver:
				log.Println("Received message in gameroom: ", m)
				for _, s := range gr.senders {
					s <- m
				}
			}
		}

	}()
}

// New creates a gameroom object.
func New() *Gameroom {
	room := &Gameroom{
		id:       generateString(10),
		senders:  []chan string{},
		receiver: make(chan string),
	}
	room.start()
	return room
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func generateString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
