package legionary

import (
	"github.com/heavycannon/heavycannon/body"
	"github.com/heavycannon/heavycannon/math/vector"
	"github.com/heavycannon/heavycannon/shape/circle"
	"github.com/hueypark/ai/ctx"
	"github.com/hueypark/ai/task/move_to_pos"
	"gitlab.com/legionary/legionary"
	"gitlab.com/legionary/legionary/composite"
)

type Legionary struct {
	id                   int64
	body                 *body.Body
	forward              vector.Vector
	acceleration         float64
	maxSpeed             float64
	holdAcceleration     float64
	maxHoldPositionSpeed float64
	bt                   legionary.BehaviorTree
	bb                   *blackboard
}

func New() *Legionary {
	l := &Legionary{}
	l.id = ctx.IdGenerator.Generate()
	l.body = createBody()
	l.forward = vector.Vector{X: 0, Y: 1}
	l.acceleration = 100
	l.maxSpeed = 100
	l.holdAcceleration = 50
	l.maxHoldPositionSpeed = 10
	ctx.PhysicsWorld.Add(l.body)

	bt := legionary.BehaviorTree{}
	l.bb = &blackboard{}
	sequence := &composite.Sequence{}
	sequence.AddChild(move_to_pos.New(l, l.bb))
	bt.SetRoot(sequence)
	l.bt = bt

	return l
}

func (l *Legionary) Update(delta float64) {
	l.bt.Update(delta)

	l.updateFriction()
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

func (l *Legionary) SetDest(dest vector.Vector) {
	l.bb.SetDest(dest)
}

func (l *Legionary) MoveTo(dest vector.Vector) {
	dir := vector.Subtract(dest, l.body.Position())
	dir.Normalize()
	l.forward = dir

	acc := vector.Multiply(l.forward, l.acceleration)

	forwardSpeed := vector.Dot(l.forward, l.body.Velocity)
	if l.maxSpeed < forwardSpeed {
		return
	}

	l.body.AddForce(acc)
}

func (l *Legionary) HoldPosition(pos vector.Vector) {
	dir := vector.Subtract(pos, l.body.Position())
	temp := dir.Size()
	dir.Normalize()

	acc := vector.Multiply(dir, l.holdAcceleration*temp)

	l.body.AddForce(acc)
}

func (l *Legionary) updateFriction() {
	backward := vector.Invert(l.forward)
	backwardSpeed := vector.Dot(backward, l.body.Velocity)

	if 30 < backwardSpeed {
		l.body.AddForce(vector.Multiply(l.forward, 100))
	}

	right := vector.Vector{X: l.forward.Y, Y: -l.forward.X}
	rightSpeed := vector.Dot(right, l.body.Velocity)
	if rightSpeed < 0 {
		l.body.AddForce(vector.Multiply(right, 50))
	} else {
		left := vector.Invert(right)
		l.body.AddForce(vector.Multiply(left, 50))
	}
}

func createBody() *body.Body {
	b := body.New()
	b.SetMass(1)
	b.SetShape(&circle.Circle{10})

	return b
}
