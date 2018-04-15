package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/heavycannon/heavycannon/math/vector"
	"github.com/hueypark/ai/actor/legionary"
	"github.com/hueypark/ai/ctx"
	"github.com/hueypark/ai/renderer"
	"github.com/hueypark/ai/world"
	"sync"
	"time"
)

var l *legionary.Legionary

func main() {
	ctx.Init()

	w := world.New()
	l = legionary.New()
	w.AddActor(l)

	lastUpdateTime := time.Now()
	var mux sync.RWMutex
	go func() {
		ticker := time.NewTicker(time.Second / 60)
		for {
			<-ticker.C

			func() {
				mux.Lock()
				defer mux.Unlock()

				delta := time.Since(lastUpdateTime).Seconds()
				lastUpdateTime = time.Now()

				ctx.PhysicsWorld.Tick(delta)

				w.Update(delta)
			}()
		}
	}()

	ebiten.Run(
		func(screen *ebiten.Image) error {
			mux.RLock()
			defer mux.RUnlock()

			handleInput(screen)

			for _, actor := range w.Actors() {
				renderer.Render(screen, actor)
			}

			return nil
		},
		ctx.ScreenWidth,
		ctx.ScreenHeight,
		1,
		"AI")
}

func handleInput(screen *ebiten.Image) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := renderer.PosToWorld(ebiten.CursorPosition())

		l.SetDest(vector.Vector{float64(x), float64(y)})
	}
}
