package ppu

import (
	"turtle/internal/chips/vram"
	"turtle/internal/pallette"
)

const (
	ScreenHeight = 128
	ScreenWidth  = 128
)

type PPU struct {
	vram         *vram.VRAM
	Transparent  pallette.Color
	Layers       map[Layer]GraphicsLayer
	currentLayer Layer
}

func New() *PPU {
	p := &PPU{
		vram:         vram.New(),
		currentLayer: SpriteLayer,
		Layers:       map[Layer]GraphicsLayer{},
	}

	p.Layers[BackgroundLayer] = newGraphicsLayer()
	p.Layers[SpriteLayer] = newGraphicsLayer()
	p.Layers[WindowLayer] = newGraphicsLayer()
	return p
}

func (p *PPU) GetFrame() []byte {
	return p.Layers[p.currentLayer].GetFrame()
}

func (p *PPU) Put(x, y uint8, c pallette.Color) {
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
	p.currentLayer = Layer(i)
}
