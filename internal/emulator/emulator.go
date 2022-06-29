package emulator

import (
	"image"
	"turtle/config"
	"turtle/internal/cart"
	"turtle/internal/chips/font"
	"turtle/internal/chips/ppu"
	"turtle/internal/vector"

	"github.com/hajimehoshi/ebiten/v2"
)

type Emulator struct {
	ppu                *ppu.PPU
	fontProcessingUnit *font.FontProcessingUnit
	Cart               cart.Cart
}

func New(cartpath string) Emulator {
	e := Emulator{
		ppu:                ppu.New(),
		fontProcessingUnit: &font.FontProcessingUnit{},
	}
	e.Cart = cart.NewCart(cartpath, vector.New(e.ppu), e.fontProcessingUnit)

	e.Cart.Init()

	return e
}

func (e Emulator) Update() {
	e.Cart.Update()
}

func (e Emulator) Draw(screen *ebiten.Image) {
	gamescreen := screen.SubImage(
		image.Rect(0, 0,
			config.Get().Window.Width,
			config.Get().Window.Height)).(*ebiten.Image)
	gamescreen.ReplacePixels(e.ppu.GetFrame())
	e.Cart.Render()
	e.ppu.Swap()
}
