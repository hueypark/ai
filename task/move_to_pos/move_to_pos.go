package move_to_pos

import (
	"github.com/heavycannon/heavycannon/math/vector"
	"gitlab.com/legionary/legionary"
)

type MoveToPos struct {
	legionary.Node
	actor actor
	dest  vector.Vector
}

type actor interface {
	Pos() vector.Vector
	MoveTo(dest vector.Vector)
}

func New(actor actor, dest vector.Vector) *MoveToPos {
	task := MoveToPos{}
	task.dest = dest
	task.actor = actor
	return &task
}

func (task *MoveToPos) Init() {
}

func (task *MoveToPos) Update(delta float64) legionary.State {
	task.actor.MoveTo(task.dest)

	if vector.Subtract(task.actor.Pos(), task.dest).SizeSquared() < 10 {
		return legionary.Success
	}

	return legionary.Running
}
