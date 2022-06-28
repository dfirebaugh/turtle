package ebitenrunner

import "github.com/hajimehoshi/ebiten/v2"

type Cart interface {
	Update()
	Render()
	Init()
}

type CartRunner struct {
	Cart        Cart
	initialized bool
}

func (cr *CartRunner) Update() {
	if !cr.initialized {
		cr.Cart.Init()
		cr.initialized = true
	}
	cr.Cart.Update()
}

func (cr CartRunner) Draw(dst *ebiten.Image) {
	cr.Cart.Render()
}
