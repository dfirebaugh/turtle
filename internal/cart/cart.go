package cart

import (
	"turtle/config"
	"turtle/internal/vm"

	lua "github.com/yuin/gopher-lua"
)

type Cart struct {
	state *lua.LState
	gp    vm.GraphicsPipeline
	lvm   vm.LuaVM
}

func NewCart(gp vm.GraphicsPipeline, fp vm.FontPipeline) *Cart {
	return &Cart{
		gp:  gp,
		lvm: vm.NewLuaVM(gp, fp),
	}
}

func (gr *Cart) LoadCart(cartCode string) error {
	config.Reset()
	state := lua.NewState()

	if err := state.DoString(cartCode); err != nil {
		return err
	}
	gr.lvm.LoadCart(state)
	gr.state = state

	gr.Init()
	return nil
}

func (gr *Cart) LoadCartFromFile(cartPath string) error {
	config.Reset()
	state := lua.NewState()

	if err := state.DoFile(cartPath); err != nil {
		return err
	}
	gr.lvm.LoadCart(state)
	gr.state = state

	gr.Init()

	return nil
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
