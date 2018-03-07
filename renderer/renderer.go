package renderer

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/heavycannon/heavycannon/math/rotator"
	"github.com/heavycannon/heavycannon/math/vector"
	"github.com/hueypark/ai/ctx"
	"golang.org/x/image/colornames"
)

var radius = 10.0

type object interface {
	Pos() vector.Vector
	Forward() vector.Vector
}

func Render(screen *ebiten.Image, o object) {
	pos := worldPosToRenderer(o.Pos())
	dir := worldDirToRenderer(o.Forward())
	renderCircle(screen, pos, dir)
}

func renderCircle(screen *ebiten.Image, pos, forward vector.Vector) {
	resolution := 32
	degrees := float64(360) / float64(resolution)
	start := vector.Add(pos, vector.Multiply(vector.Vector{X: 0, Y: -1}, radius))
	for i := 0; i <= resolution; i++ {
		r := rotator.Rotator{Degrees: degrees * float64(i)}
		end := vector.Add(pos, r.RotateVector(vector.Multiply(vector.Vector{X: 0, Y: -1}, radius)))
		ebitenutil.DrawLine(screen, start.X, start.Y, end.X, end.Y, colornames.Green)
		start = end
	}

	end := vector.Add(pos, vector.Multiply(forward, radius))
	ebitenutil.DrawLine(screen, pos.X, pos.Y, end.X, end.Y, colornames.Green)
}

func worldPosToRenderer(pos vector.Vector) vector.Vector {
	return vector.Vector{
		X: pos.X + float64(ctx.ScreenWidth)*0.5,
		Y: -pos.Y + float64(ctx.ScreenHeight)*0.5}
}

func worldDirToRenderer(dir vector.Vector) vector.Vector {
	return vector.Vector{X: dir.X, Y: -dir.Y}
}
