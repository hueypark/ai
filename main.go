package main

import (
	"time"

	"github.com/hajimehoshi/ebiten"

	"github.com/hueypark/ai/actor/legionary"
	"github.com/hueypark/ai/world"
)

func main() {
	w := world.New()
	l := legionary.New()
	w.AddActor(l)

	lastUpdateTime := time.Now()

	ebiten.Run(
		func(screen *ebiten.Image) error {
			delta := time.Since(lastUpdateTime)
			lastUpdateTime = time.Now()

			w.Update(delta)

			for _, actor := range w.Actors() {
				actor.Render(screen)
			}

			return nil
		},
		1024,
		768,
		1,
		"AI")
}
