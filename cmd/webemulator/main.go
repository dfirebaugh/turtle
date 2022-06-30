// build: GOOS=js GOARCH=wasm go build -o main.wasm main.go

package main

import (
	"syscall/js"
	"turtle/config"
	"turtle/internal/emulator"
	"turtle/internal/emulator/engine/ebitenrunner"
	"turtle/internal/emulator/engine/ebitenrunner/ebitenwrapper"

	"golang.org/x/image/colornames"
)

var cartCode string

func loadCart(s string) {
	println(s)
	// document := js.Global().Get("document")
	// p := document.Call("createElement", "p")
	// p.Set("innerHTML", message)
	// document.Get("body").Call("appendChild", p)
}

func setJSFuncs() {
	c := make(chan bool)
	js.Global().Set("loadCart", js.FuncOf(loadCart))
	<-c
}

func main() {
	runner := emulator.New(cartCode)

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
