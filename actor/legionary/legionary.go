package legionary

import (
	"github.com/heavycannon/heavycannon/body"
	"github.com/heavycannon/heavycannon/math/vector"
	"github.com/heavycannon/heavycannon/shape/circle"
	"github.com/hueypark/ai/ctx"
	"time"
)

type Legionary struct {
	id      int64
	body    *body.Body
	forward vector.Vector
}

func New() *Legionary {
	l := &Legionary{}
	l.id = ctx.IdGenerator.Generate()
	l.body = createBody()
	ctx.PhysicsWorld.Add(l.body)

	return l
}

var temp = true
var xxTime time.Duration

func (l *Legionary) Update(delta time.Duration) {
}

func (l *Legionary) ID() int64 {
	return l.id
}

func (l *Legionary) Pos() vector.Vector {
	return l.body.Position()
}

func (l *Legionary) Forward() vector.Vector {
	return l.forward
}

func createBody() *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(&circle.Circle{10})

	return b
}
