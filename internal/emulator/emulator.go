package emulator

import (
	"image"
	"turtle/internal/cart"
	"turtle/internal/emulator/chips/font"
	"turtle/internal/emulator/chips/ppu"
	"turtle/internal/emulator/engine"
)

type gameEngine interface {
	Run()
	Reset(interface{})
}

type Emulator struct {
	pixelProcessingUnit *ppu.PPU
	fontProcessingUnit  *font.FontProcessingUnit
	Cart                *cart.Cart
	engine              gameEngine
}

func New() Emulator {
	e := Emulator{
		pixelProcessingUnit: ppu.New(),
		fontProcessingUnit:  &font.FontProcessingUnit{},
	}
	e.Cart = cart.NewCart(e.pixelProcessingUnit, e.fontProcessingUnit)
	e.Cart.Init()
	e.engine = engine.New(e)

	return e
}

func (e Emulator) LoadCart(c string) error {
	err := e.Cart.LoadCart(c)
	e.engine.Reset(e)

	return err
}

func (e Emulator) LoadCartFromFile(c string) {
	e.Cart.LoadCartFromFile(c)
	e.engine.Reset(e)
}

func (e Emulator) Run() {
	e.engine.Run()
}

func (e Emulator) Update() {
	if e.Cart == nil {
		println("no cart")
		return
	}
	e.Cart.Update()
}

func (e Emulator) Render(screen *image.RGBA) {
	if e.Cart == nil {
		return
	}
	screen.Pix = e.pixelProcessingUnit.GetFrame()
	e.Cart.Render()
	e.pixelProcessingUnit.Swap()
}
