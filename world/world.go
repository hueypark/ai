package world

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	"log"
)

type World struct {
	actors map[int64]actor
}

type actor interface {
	ID() int64
	Render(screen *ebiten.Image)
}

func New() *World {
	w := &World{}
	w.actors = make(map[int64]actor)

	return w
}

func (w *World) Update(delta time.Duration) error {
	log.Println(delta)

	return nil
}

func (w *World) AddActor(a actor) {
	w.actors[a.ID()] = a
}

func (w *World) Actors() map[int64]actor {
	return w.actors
}
