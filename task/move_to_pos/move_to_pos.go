package move_to_pos

import (
	"github.com/heavycannon/heavycannon/math/vector"
	"gitlab.com/legionary/legionary"
)

type MoveToPos struct {
	legionary.Node
	actor actor
	bb    tmep
	dest  vector.Vector
}

type actor interface {
	Pos() vector.Vector
	MoveTo(dest vector.Vector)
}

type tmep interface {
	Dest() vector.Vector
}

func New(actor actor, bb tmep) *MoveToPos {
	task := MoveToPos{}
	task.actor = actor
	task.bb = bb
	return &task
}

func (task *MoveToPos) Init() {
}

func (task *MoveToPos) Update(delta float64) legionary.State {
	task.actor.MoveTo(task.bb.Dest())

	if vector.Subtract(task.actor.Pos(), task.dest).SizeSquared() < 400 {
		return legionary.Success
	}

	return legionary.Running
}
