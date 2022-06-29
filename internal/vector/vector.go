package vector

import (
	"math"
	"turtle/internal/pallette"
	"turtle/pkg/gamemath"
)

type memory interface {
	Put(x uint8, y uint8, c pallette.Color)
	Clear()
}

type Vector struct {
	memory memory
}

func New(memory memory) Vector {
	return Vector{
		memory: memory,
	}
}

func (v Vector) Rect(rect gamemath.Rect, color pallette.Color) {
	x0 := rect[0]
	y0 := rect[1]
	y1 := rect[3] + y0

	for i := 0; i < int(rect[2]); i++ {
		v.Line(gamemath.MakeVector(x0+float64(i), y0), gamemath.MakeVector(x0+float64(i), y1), 1, color)
	}
}
func (v Vector) Line(v0 gamemath.Vector, v1 gamemath.Vector, stroke float64, color pallette.Color) {
	x := v0[0]
	y := v0[1]

	heading := v1.GetHeading(v0)

	dx := math.Cos(heading)
	dy := math.Sin(heading)

	for i := 0; float64(i) < gamemath.GetDistance(v0, v1); i++ {
		v.memory.Put(uint8(x+(float64(i)*dx)), uint8(y+(float64(i)*dy)), color)
	}
}
func (v Vector) Circ(c gamemath.Circle, color pallette.Color) {
	for i := 0; i < 360; i++ {
		x1 := c.R * math.Cos(float64(i)*math.Pi/180)
		y1 := c.R * math.Sin(float64(i)*math.Pi/180)

		v.memory.Put(uint8(c.X+x1+c.R), uint8(c.Y+y1+c.R), color)
	}
}
func (v Vector) Tri(a, b, c uint8) {

}

func (v Vector) Clear() {
	v.memory.Clear()
}
