package cart

import (
	"turtle/internal/vm"

	lua "github.com/yuin/gopher-lua"
)

type Cart struct {
	state *lua.LState
	gp    vm.GraphicsPipeline
	lvm   vm.LuaVM
}

func NewCart(cartpath string, gp vm.GraphicsPipeline, fp vm.FontPipeline) Cart {
	state := lua.NewState()
	if err := state.DoFile(cartpath); err != nil {
		panic(err)
	}
	lvm := vm.NewLuaVM(gp, fp, state)

	return Cart{
		state: state,
		gp:    gp,
		lvm:   lvm,
	}
}

func (gr Cart) Init() {
	if gr.state == nil {
		return
	}
	gr.lvm.Init(gr.state)
}

func (gr Cart) Update() {
	if gr.state == nil {
		return
	}
	gr.lvm.UpdateCalls(gr.state)
}

func (gr Cart) Render() {
	if gr.state == nil {
		return
	}
	gr.lvm.DrawCalls(gr.state)
}
