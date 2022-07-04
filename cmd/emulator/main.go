package main

import (
	"flag"
	"turtle/internal/emulator"
)

var cartPath string

func init() {
	flag.StringVar(&cartPath, "cart", "./examples/raycast.lua", "relative path to cart file")
}

func main() {
	flag.Parse()
	println(cartPath)

	em := emulator.New()
	em.LoadCartFromFile(cartPath)
	em.Run()
}
