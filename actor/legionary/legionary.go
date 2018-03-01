package legionary

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/hueypark/ai/ctx"
	"log"
)

type Legionary struct {
	id int64
}

func New() *Legionary {
	l := &Legionary{}
	l.id = ctx.IdGenerator.Generate()

	return l
}

func (l *Legionary) ID() int64 {
	return l.id
}

func (l *Legionary) Render(screen *ebiten.Image) {
	log.Println("I am legionary")
}
