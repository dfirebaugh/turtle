package ppu

import (
	"turtle/internal/emulator/chips/math"
	"turtle/internal/emulator/chips/ppu/plotter"
	"turtle/internal/emulator/chips/vram"
)

type GraphicsPipeline interface {
	Rect(rect math.Rect, color uint8)
	Line(v0 math.Vector, v1 math.Vector, color uint8)
	Triangle(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8)
	Circle(circle math.Circle, color uint8)
	Point(x uint16, y uint16, color uint8)
	ShiftLayer(i uint8)
	Clear()
	RenderSprite(sprite []uint8, x, y float64)
}

type PPU struct {
	plotter      plotter.Plotter
	vram         *vram.VRAM
	Transparent  uint8
	Layers       map[Layer]GraphicsLayer
	currentLayer Layer
}

func New(v *vram.VRAM) *PPU {
	p := &PPU{
		vram:         v,
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

func (p *PPU) Put(x, y uint16, c uint8) {
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
func (p PPU) Circle(circle math.Circle, color uint8) {
	p.plotter.Circle(circle, color)
}
func (p PPU) CircleFill(circle math.Circle, color uint8) {
	p.plotter.Circle(circle, color)
}
func (p PPU) Point(x uint16, y uint16, color uint8) {
	p.plotter.Point(x, y, color)
}
func (p PPU) RenderSprite(sprite []uint8, x, y float64) {
	p.plotter.RenderSprite(sprite, x, y)
}
