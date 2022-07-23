package plotter

import (
	"image/color"
	"turtle/internal/emulator/chips/math"
	"turtle/internal/emulator/chips/ppu/pallette"

	"golang.org/x/image/colornames"
	"tinygo.org/x/tinydraw"
)

type memory interface {
	Put(x uint16, y uint16, c uint8)
	SetPixel(x, y int16, c color.RGBA)
	Display() error
	Size() (int16, int16)
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

func (p Plotter) Point(x uint16, y uint16, color uint8) {
	p.memory.Put(x, y, color)
}

func (p Plotter) Rect(rect math.Rect, clr uint8) {
	color, ok := pallette.GetColor(clr).(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.FilledRectangle(p.memory, int16(rect[0]), int16(rect[1]), int16(rect[2]), int16(rect[3]), color)
}

func (p Plotter) Line(v0 math.Vector, v1 math.Vector, clr uint8) {
	color, ok := pallette.GetColor(clr).(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.Line(p.memory, int16(v0[0]), int16(v0[1]), int16(v1[0]), int16(v1[1]), color)
}

func (p Plotter) Triangle(v0 math.Vector, v1 math.Vector, v2 math.Vector, clr uint8) {
	color, ok := pallette.GetColor(clr).(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.FilledTriangle(p.memory, int16(v0[0]), int16(v0[1]), int16(v1[0]), int16(v1[1]), int16(v2[0]), int16(v2[1]), color)
}

func (p Plotter) Circle(c math.Circle, clr uint8) {
	color, ok := pallette.GetColor(clr).(color.RGBA)
	if !ok {
		color = colornames.Black
	}
	tinydraw.FilledCircle(p.memory, int16(c.X), int16(c.Y), int16(c.R*2), color)
}

func (p Plotter) RenderSprite(sprite []uint8, x, y float64) {
	for i, color := range sprite {
		p.Point(uint16(i%8)+uint16(x), uint16(i/8)+uint16(y), color)
	}
}

func (p Plotter) Clear() {
	p.memory.Clear()
}
