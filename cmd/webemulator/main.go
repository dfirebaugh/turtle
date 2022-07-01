//go:build (js && ignore) || wasm
// +build js,ignore wasm

package main

import (
	"syscall/js"
	"turtle/internal/emulator"
	"turtle/internal/emulator/engine/ebitenrunner/ebitenwrapper"
)

var cartCode string
var game *ebitenwrapper.Game

func loadCart(value js.Value, args []js.Value) interface{} {
	codeText := args[0].String()
	if codeText == "" {
		println("no code submitted...")
		return nil
	}
	runner := emulator.New()
	runner.Cart.LoadCart(codeText)
	game.Reset(runner)
	return nil
}

func setJSFuncs() {
	js.Global().Set("loadCart", js.FuncOf(loadCart))
}

func main() {
	setJSFuncs()
	game = ebitenwrapper.New()
	game.Run()
}
