package ppu

import (
	"turtle/internal/emulator/chips/math"
	"turtle/internal/emulator/chips/ppu/plotter"
	"turtle/internal/emulator/chips/vram"
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
	plotter      plotter.Plotter
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
	p.plotter = plotter.New(p.vram)

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
	p.plotter.Rect(rect, color)
}
func (p PPU) Line(v0 math.Vector, v1 math.Vector, color uint8) {
	p.plotter.Line(v0, v1, color)
}
func (p PPU) Triangle(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8) {
	p.plotter.Triangle(v0, v1, v2, color)
}
func (p PPU) TriangleFill(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8) {
	p.plotter.Triangle(v0, v1, v2, color)
}
func (p PPU) Circ(circle math.Circle, color uint8) {
	p.plotter.Circ(circle, color)
}
func (p PPU) CircFill(circle math.Circle, color uint8) {
	p.plotter.Circ(circle, color)
}
func (p PPU) Point(x uint8, y uint8, color uint8) {
	p.plotter.Point(x, y, color)
}
