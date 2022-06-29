package pixelglrunner

import "github.com/faiface/pixel/pixelgl"

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

func (cr CartRunner) Draw(*pixelgl.Window) {
	cr.Cart.Render()
}
