package main

import (
	"flag"
	"turtle/config"
	"turtle/internal/emulator"
	"turtle/internal/emulator/engine/ebitenrunner"
	"turtle/internal/emulator/engine/ebitenrunner/ebitenwrapper"

	"golang.org/x/image/colornames"
)

var cartPath string

func init() {
	flag.StringVar(&cartPath, "cart", "main.lua", "relative path to cart file")
}

func main() {
	flag.Parse()

	runner := emulator.New(cartPath)
	systems := []ebitenrunner.System{runner}
	drawables := []ebitenrunner.Drawable{runner}

	c := config.Get()

	e := &ebitenwrapper.Game{
		Scene:           ebitenrunner.New(systems, drawables),
		WindowTitle:     c.Title,
		WindowScale:     c.ScaleFactor,
		Width:           c.Window.Width,
		Height:          c.Window.Height,
		BackgroundColor: colornames.Skyblue,
	}

	e.Run()
}
