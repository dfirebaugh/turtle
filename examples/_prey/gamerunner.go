package main

import (
	_ "embed"
	"turtle/internal/vm"

	"github.com/hajimehoshi/ebiten/v2"
	lua "github.com/yuin/gopher-lua"
)

//go:embed main.lua
var LuaScript string

type gameRunner struct {
	state *lua.LState
}

func NewGameRunner() gameRunner {
	state := lua.NewState()
	if err := state.DoString(LuaScript); err != nil {
		panic(err)
	}
	state.SetGlobal("rect", state.NewFunction(vm.MakeRect(f)))
	state.SetGlobal("clear", state.NewFunction(vm.Clear(f)))
	vm.LoadGlobals(state)

	return gameRunner{
		state: state,
	}
}

func (gr gameRunner) Update() {
	if gr.state == nil {
		return
	}
	vm.UpdateCalls(gr.state)
}

func (gr gameRunner) Draw(screen *ebiten.Image) {
	if gr.state == nil {
		return
	}
	vm.DrawCalls(gr.state)
}
