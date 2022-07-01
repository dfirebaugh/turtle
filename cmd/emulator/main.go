package main

import (
	"flag"
	"turtle/internal/emulator"
	"turtle/internal/emulator/engine/ebitenrunner/ebitenwrapper"
)

var cartPath string

func init() {
	flag.StringVar(&cartPath, "cart", "./examples/raycast/raycast.lua", "relative path to cart file")
}

func main() {
	flag.Parse()

	runner := emulator.New()
	e := ebitenwrapper.New()
	runner.Cart.LoadCartFromFile(cartPath)
	e.Scene.Reset(runner)

	e.Run()
}
