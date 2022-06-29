package main

import (
	"flag"
	"turtle/pkg/turtle"
)

var cartPath string

func init() {
	flag.StringVar(&cartPath, "cart", "main.lua", "relative path to cart file")
}

func main() {
	flag.Parse()

	println(cartPath)

	turtle.RunCart(cartPath)
}
