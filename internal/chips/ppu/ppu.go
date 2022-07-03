package ppu

import (
	"turtle/internal/chips/math"
	"turtle/internal/chips/ppu/vector"
	"turtle/internal/chips/vram"
)

const (
	ScreenHeight = 128
	ScreenWidth  = 128
)

type GraphicsPipeline interface {
	Rect(rect math.Rect, color uint8)
	Line(v0 math.Vector, v1 math.Vector, color uint8)
	Triangle(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8)
	Circ(circle math.Circle, color uint8)
	Point(x uint8, y uint8, color uint8)
	ShiftLayer(i uint8)
	Clear()
}

type PPU struct {
	vector       vector.Vector
	vram         *vram.VRAM
	Transparent  uint8
	Layers       map[Layer]GraphicsLayer
	currentLayer Layer
}

func New() *PPU {
	p := &PPU{
		vram:         vram.New(),
		currentLayer: SpriteLayer,
		Layers:       map[Layer]GraphicsLayer{},
	}
	p.vector = vector.New(p.vram)

	p.Layers[BackgroundLayer] = newGraphicsLayer()
	p.Layers[SpriteLayer] = newGraphicsLayer()
	p.Layers[WindowLayer] = newGraphicsLayer()
	return p
}

func (p *PPU) GetFrame() []byte {
	return p.Layers[p.currentLayer].GetFrame()
}

func (p *PPU) Put(x, y uint8, c uint8) {
	p.vram.Put(x, y, c)
}

func (p *PPU) Swap() {
	p.vram.Swap()
	p.Layers[p.currentLayer] = GraphicsLayer(p.vram.GetBuffer())
}

func (p *PPU) Clear() {
	p.vram.Clear()
}

func (p *PPU) ShiftLayer(i uint8) {
	p.currentLayer = Layer(i % 3)
}

func (p PPU) Rect(rect math.Rect, color uint8) {
	p.vector.Rect(rect, color)
}
func (p PPU) Line(v0 math.Vector, v1 math.Vector, color uint8) {
	p.vector.Line(v0, v1, color)
}
func (p PPU) Triangle(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8) {
	p.vector.Triangle(v0, v1, v2, color)
}
func (p PPU) Circ(circle math.Circle, color uint8) {
	p.vector.Circ(circle, color)
}
func (p PPU) Point(x uint8, y uint8, color uint8) {
	p.vector.Point(x, y, color)
}
