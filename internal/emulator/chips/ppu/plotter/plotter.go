package plotter

import "turtle/internal/emulator/chips/math"

type memory interface {
	Put(x uint8, y uint8, c uint8)
	Clear()
}

type Plotter struct {
	memory memory
}

func New(memory memory) Plotter {
	return Plotter{
		memory: memory,
	}
}

func (p Plotter) Point(x uint8, y uint8, color uint8) {
	p.memory.Put(x, y, color)
}

func (p Plotter) Rect(rect math.Rect, color uint8) {
	x0 := rect[0]
	y0 := rect[1]
	y1 := rect[3] + y0

	for i := 0; i < int(rect[2]); i++ {
		p.Line(math.MakeVector(x0+float64(i), y0), math.MakeVector(x0+float64(i), y1), color)
	}
}
func (p Plotter) Line(v0 math.Vector, v1 math.Vector, color uint8) {
	x := v0[0]
	y := v0[1]

	heading := v1.GetHeading(v0)

	dx := math.Cos(heading)
	dy := math.Sin(heading)

	for i := 0; float64(i) < math.GetDistance(v0, v1); i++ {
		p.memory.Put(uint8(x+(float64(i)*dx)), uint8(y+(float64(i)*dy)), color)
	}
}

func (p Plotter) FillBottomTri(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8) {
	invslope1 := (v1[0] - v0[0]) / (v1[1] - v0[1])
	invslope2 := (v2[0] - v0[0]) / (v2[1] - v0[1])

	curX1 := v0[0]
	curX2 := v0[0]

	for scanLineY := v0[1]; scanLineY <= v1[1]; scanLineY++ {
		p.Line(math.MakeVector(curX1, scanLineY), math.MakeVector(curX2, scanLineY), color)
		curX1 += invslope1
		curX2 += invslope2
	}
}
func (p Plotter) FillTopTri(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8) {
	invslope1 := (v2[0] - v0[0]) / (v2[1] - v0[1])
	invslope2 := (v2[0] - v1[0]) / (v2[1] - v1[1])

	curX1 := v2[0]
	curX2 := v2[0]

	for scanLineY := v2[1]; scanLineY > v0[1]; scanLineY-- {
		p.Line(math.MakeVector(curX1, scanLineY), math.MakeVector(curX2, scanLineY), color)
		curX1 -= invslope1
		curX2 -= invslope2
	}
}
func (p Plotter) Triangle(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8) {
	p.Line(v0, v1, color)
	p.Line(v0, v2, color)
	p.Line(v1, v2, color)

	p.FillBottomTri(v0, v1, v2, color)
	p.FillTopTri(v0, v1, v2, color)
}

func (p Plotter) CircleOutline(c math.Circle, color uint8) {
	for i := 0; i < 360; i++ {
		x1 := c.R * math.Cos(float64(i)*math.Pi/180)
		y1 := c.R * math.Sin(float64(i)*math.Pi/180)

		p.memory.Put(uint8(c.X+x1+c.R), uint8(c.Y+y1+c.R), color)
	}
}
func (p Plotter) Circ(c math.Circle, color uint8) {
	p.CircleDumbFill(c, color)
	p.CircleOutline(c, color)
}

func (p Plotter) CircleFill(c math.Circle, color uint8) {

}

// DumbFill draws a rect inside the circle
// it doesn't fit though :3
func (p Plotter) CircleDumbFill(c math.Circle, color uint8) {
	p.Rect(math.MakeRect(c.X+2, c.Y+2, c.R*2-3, c.R*2-3), color)
}

func (p Plotter) Clear() {
	p.memory.Clear()
}
