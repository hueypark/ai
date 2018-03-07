package ctx

import (
	"github.com/heavycannon/heavycannon"
	"github.com/heavycannon/heavycannon/math/vector"
	"github.com/hueypark/framework/core/id_generator"
)

var (
	IdGenerator  id_generator.IdGenerator
	PhysicsWorld *heavycannon.World
	ScreenWidth  = 1024
	ScreenHeight = 768
)

func Init() {
	PhysicsWorld = heavycannon.New()
	PhysicsWorld.SetGravity(vector.ZERO())
}
