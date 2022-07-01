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
	err := runner.Cart.LoadCartFromFile(cartPath)
	if err != nil {
		panic(err)
	}
	e.Scene.Reset(runner)

	e.Run()
}
