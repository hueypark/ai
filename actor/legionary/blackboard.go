package legionary

import "github.com/heavycannon/heavycannon/math/vector"

type blackboard struct {
	dest vector.Vector
}

func (bb *blackboard) SetDest(dest vector.Vector) {
	bb.dest = dest
}

func (bb blackboard) Dest() vector.Vector {
	return bb.dest
}
