package world

import (
	"time"

	"github.com/heavycannon/heavycannon/math/vector"
)

type World struct {
	actors map[int64]actor
}

type actor interface {
	ID() int64
	Pos() vector.Vector
	Forward() vector.Vector
	Update(time.Duration)
}

func New() *World {
	w := &World{}
	w.actors = make(map[int64]actor)

	return w
}

func (w *World) Update(delta time.Duration) error {
	for _, actor := range w.actors {
		actor.Update(delta)
	}

	return nil
}

func (w *World) AddActor(a actor) {
	w.actors[a.ID()] = a
}

func (w *World) Actors() map[int64]actor {
	return w.actors
}
