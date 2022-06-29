package vm

import (
	"math/rand"
	"turtle/config"
	"turtle/internal/graphics"
	"turtle/internal/pallette"
	"turtle/pkg/gamemath"

	lua "github.com/yuin/gopher-lua"
)

type LuaVM struct {
	gp      *graphics.GraphicsPipeline
	globals map[string]lua.LGFunction
}

func NewLuaVM(gp *graphics.GraphicsPipeline) LuaVM {
	lvm := LuaVM{
		gp: gp,
	}
	lvm.globals = map[string]lua.LGFunction{
		"SCREENH":     lvm.getScreenHeight,
		"SCREENW":     lvm.getScreenWidth,
		"RND":         lvm.random,
		"SCALEFACTOR": lvm.scaleFactor,
	}
	return lvm
}

func (l LuaVM) LoadGlobals(L *lua.LState) {
	for key, fn := range l.globals {
		L.SetGlobal(key, L.NewFunction(fn))
	}
}

func (LuaVM) getScreenHeight(L *lua.LState) int {
	L.Push(lua.LNumber(config.Get().Window.Height))
	return 1
}
func (LuaVM) getScreenWidth(L *lua.LState) int {
	L.Push(lua.LNumber(config.Get().Window.Height))
	return 1
}

func (l LuaVM) Clear(L *lua.LState) int {
	l.gp.Clear()
	return 0
}

func (l LuaVM) GetRandomColor(L *lua.LState) int {
	return 1
}

func (l LuaVM) PrintAt(L *lua.LState) int {
	// l.gp.Print()
	return 0
}

func (l LuaVM) MakeRect(L *lua.LState) int {
	if l.gp == nil {
		return 0
	}
	x := float64(L.ToNumber(1))
	y := float64(L.ToNumber(2))
	w := float64(L.ToNumber(3))
	h := float64(L.ToNumber(4))
	color := int(L.ToNumber(5))
	// l.gp.Rect(gamemath.MakeRect(x-1, y-1, w+2, h+2), 1)
	l.gp.Rect(gamemath.MakeRect(x, y, w, h), pallette.Color(color))
	return 0
}

func (l LuaVM) MakeLine(L *lua.LState) int {
	x0 := float64(L.ToNumber(1))
	y0 := float64(L.ToNumber(2))
	x1 := float64(L.ToNumber(3))
	y1 := float64(L.ToNumber(4))

	c := pallette.Color(L.ToNumber(5))
	l.gp.Line(gamemath.MakeVector(x0, y0), gamemath.MakeVector(x1, y1), 1, c)
	return 0
}

func (l LuaVM) MakeCircle(L *lua.LState) int {
	if l.gp == nil {
		return 0
	}
	x := float64(L.ToNumber(1))
	y := float64(L.ToNumber(2))
	r := float64(L.ToNumber(3))
	color := int(L.ToNumber(4))
	l.gp.Circ(gamemath.MakeCircle(x, y, r), pallette.Color(color))
	return 0
}
func (LuaVM) Init(state *lua.LState) {
	if err := state.CallByParam(lua.P{
		Fn:      state.GetGlobal("INIT"), // name of Lua function
		NRet:    0,                       // number of returned values
		Protect: true,                    // return err or panic
	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
		panic(err)
	}
}
func (LuaVM) UpdateCalls(state *lua.LState) {
	if err := state.CallByParam(lua.P{
		Fn:      state.GetGlobal("UPDATE"), // name of Lua function
		NRet:    0,                         // number of returned values
		Protect: true,                      // return err or panic
	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
		panic(err)
	}
}
func (LuaVM) DrawCalls(state *lua.LState) {
	if err := state.CallByParam(lua.P{
		Fn:      state.GetGlobal("RENDER"), // name of Lua function
		NRet:    0,                         // number of returned values
		Protect: true,                      // return err or panic
	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
		panic(err)
	}
}

func (LuaVM) random(state *lua.LState) int {
	state.Push(lua.LNumber(rand.Intn(state.ToInt(1))))
	return 1
}

func (LuaVM) scaleFactor(state *lua.LState) int {
	state.Push(lua.LNumber(10))
	return 1
}
