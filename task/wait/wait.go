package wait

import (
	"github.com/heavycannon/heavycannon/math/vector"
	"gitlab.com/legionary/legionary"
)

type Wait struct {
	legionary.Node
	actor      actor
	time       float64
	remainTime float64
	waitPos    vector.Vector
}

type actor interface {
	Pos() vector.Vector
	HoldPosition(pos vector.Vector)
}

func New(actor actor, time float64) *Wait {
	task := Wait{}
	task.time = time
	task.actor = actor
	return &task
}

func (task *Wait) Init() {
	task.remainTime = task.time
	task.waitPos = task.actor.Pos()
}

func (task *Wait) Update(delta float64) legionary.State {
	task.actor.HoldPosition(task.waitPos)

	task.remainTime -= delta
	if task.remainTime < 0 {
		return legionary.Success
	}

	return legionary.Running
}
