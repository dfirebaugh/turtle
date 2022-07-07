//go:build (js && ignore) || wasm
// +build js,ignore wasm

package main

import (
	"syscall/js"
	"turtle/internal/emulator"
)

var (
	cartCode string
	game     = emulator.New()
)

func showErrorMessage(msg string) {
	js.Global().Call("showError", msg)
}

var setEditor = func(code string) {
	js.Global().Call("setEditorValue", code)
}

func loadCart(value js.Value, args []js.Value) interface{} {
	codeText := args[0].String()
	if codeText == "" {
		println("no code submitted...")
		return nil
	}
	err := game.LoadCart(codeText)

	if err != nil {
		println(err)
		showErrorMessage(err.Error())
		return nil
	}
	return nil
}

func setJSFuncs() {
	js.Global().Set("loadCart", js.FuncOf(loadCart))
}

func main() {
	setJSFuncs()
	game.SetEditorCb(setEditor)
	game.LoadCart(`function INIT()
	end
	function UPDATE()
	end
	function RENDER() 
		RECT(0, 0, 128, 128, 0)
	end
	`)
	game.Run()
}
