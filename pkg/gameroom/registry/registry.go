package registry

import (
	"errors"

	"github.com/huntsman90/litserver/pkg/gameroom"
)

// Registry is an in-memory store for active game rooms.
type Registry struct {
	rooms map[string]*gameroom.Gameroom
}

// Register a gameroom in the registry.
func (r *Registry) Register(room *gameroom.Gameroom) error {
	if _, ok := r.rooms[room.ID()]; ok {
		return errors.New("room already exists")
	}
	r.rooms[room.ID()] = room
	return nil
}

// Unregister a gameroom. This will effectively delete a game.
func (r *Registry) Unregister(room *gameroom.Gameroom) {
	if _, ok := r.rooms[room.ID()]; ok {
		delete(r.rooms, room.ID())
	}
}

// Room will fetch a gameroom using the game ID as key.
func (r *Registry) Room(key string) (*gameroom.Gameroom, error) {
	if room, ok := r.rooms[key]; ok {
		return room, nil
	}
	return nil, errors.New("room not found")
}

// New creates a gameroom registry.
func New() *Registry {
	return &Registry{
		rooms: make(map[string]*gameroom.Gameroom),
	}
}
