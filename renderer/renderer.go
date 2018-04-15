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
	pos := posFromWorld(o.Pos())
	dir := dirFromWorld(o.Forward())
	renderCircle(screen, pos, dir)
}

func PosToWorld(x, y int) (int, int) {
	return x - ctx.ScreenWidth/2, -(y - ctx.ScreenHeight/2)
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

func posFromWorld(pos vector.Vector) vector.Vector {
	return vector.Vector{
		X: pos.X + float64(ctx.ScreenWidth)*0.5,
		Y: -pos.Y + float64(ctx.ScreenHeight)*0.5}
}

func dirFromWorld(dir vector.Vector) vector.Vector {
	return vector.Vector{X: dir.X, Y: -dir.Y}
}
