package vm

import (
	"math"
	"math/rand"
	"time"
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

func NewLuaVM(gp *graphics.GraphicsPipeline, L *lua.LState) LuaVM {
	lvm := LuaVM{
		gp: gp,
	}
	lvm.setGlobals(L)

	return lvm
}

func (lvm LuaVM) setGlobals(L *lua.LState) {
	globals := map[string]lua.LGFunction{
		"SCREENH":     lvm.getScreenHeight,
		"SCREENW":     lvm.getScreenWidth,
		"RND":         lvm.random,
		"SCALEFACTOR": lvm.scaleFactor,
		"RECT":        lvm.MakeRect,
		"CIR":         lvm.MakeCircle,
		"LINE":        lvm.MakeLine,
		"CLS":         lvm.Clear,
		"COS":         lvm.Cos,
		"SIN":         lvm.Sin,
		"HEADING":     lvm.GetHeading,
	}
	for key, fn := range globals {
		L.SetGlobal(key, L.NewFunction(fn))
	}
	// L.SetGlobal("RECT", L.NewFunction(lvm.MakeRect))
	// L.SetGlobal("CIR", L.NewFunction(lvm.MakeCircle))
	// L.SetGlobal("LINE", L.NewFunction(lvm.MakeLine))
	// L.SetGlobal("CLS", L.NewFunction(lvm.Clear))
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

func (LuaVM) random(state *lua.LState) int {
	rand.Seed(time.Now().UnixNano())
	state.Push(lua.LNumber(rand.Intn(state.ToInt(1))))
	return 1
}

func (LuaVM) scaleFactor(state *lua.LState) int {
	state.Push(lua.LNumber(10))
	return 1
}

func (LuaVM) Cos(state *lua.LState) int {
	n := float64(state.ToNumber(1))

	state.Push(lua.LNumber(math.Cos(n)))
	return 1
}

func (LuaVM) GetHeading(state *lua.LState) int {
	v0 := gamemath.MakeVector(float64(state.ToNumber(1)), float64(state.ToNumber(2)))
	v1 := gamemath.MakeVector(float64(state.ToNumber(3)), float64(state.ToNumber(4)))
	state.Push(lua.LNumber(v0.GetHeading(v1)))
	return 1
}

func (LuaVM) Sin(state *lua.LState) int {
	n := float64(state.ToNumber(1))

	state.Push(lua.LNumber(math.Sin(n)))
	return 1
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
