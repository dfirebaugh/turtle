package main

import (
	"image/color"
	"turtle/config"
	"turtle/examples/prey/carts/prey"
	"turtle/internal/fb"
	"turtle/internal/govm"
	"turtle/internal/scenebuilder"
	"turtle/pkg/ebitenwrapper"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.ErrorLevel)
	frameBuffer := &fb.FrameBuffer{}

	c := config.Get()

	// gr := NewGameRunner(frameBuffer)
	cart := govm.NewCart(prey.Cart{}, frameBuffer)

	systems := []scenebuilder.System{
		cart,
	}

	drawables := []scenebuilder.Drawable{
		frameBuffer,
		cart,
	}

	game := &ebitenwrapper.Game{
		Scene:           scenebuilder.New(systems, drawables, func() {}),
		WindowTitle:     c.Title,
		WindowScale:     c.ScaleFactor,
		Width:           c.Window.Width,
		Height:          c.Window.Height,
		BackgroundColor: color.Black,
	}

	game.Run()
}
