package vm

import (
	"math/rand"
	"time"
	"turtle/config"
	"turtle/internal/chips/math"
	"turtle/internal/gamepad"

	"github.com/google/uuid"
	lua "github.com/yuin/gopher-lua"
)

type GraphicsPipeline interface {
	Rect(rect math.Rect, color uint8)
	Line(v0 math.Vector, v1 math.Vector, color uint8)
	Triangle(v0 math.Vector, v1 math.Vector, v2 math.Vector, color uint8)
	Circ(circle math.Circle, color uint8)
	Point(x uint8, y uint8, color uint8)
	ShiftLayer(i uint8)
	Clear()
}

type FontPipeline interface {
	PrintAt(string, int, int)
}

type LuaVM struct {
	gp          GraphicsPipeline
	fp          FontPipeline
	controllers []gamepad.GamePad
	globals     map[string]lua.LGFunction
	tick        *uint
}

// world clock
var tick uint = 1

func NewLuaVM(gp GraphicsPipeline, fp FontPipeline) LuaVM {
	lvm := LuaVM{
		gp: gp,
		fp: fp,
		controllers: []gamepad.GamePad{{
			Buttons: make(map[gamepad.Button]bool),
			Device:  gamepad.Keyboard{},
		}},
		tick: &tick,
	}
	lvm.initializeTick()

	return lvm
}

func (lvm LuaVM) LoadCart(state *lua.LState) {
	lvm.setGlobals(state)
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
		"TRI":         lvm.MakeTriangle,
		"POINT":       lvm.MakePoint,
		"CLS":         lvm.Clear,
		"CLR":         lvm.Clear,
		"COS":         lvm.Cos,
		"SIN":         lvm.Sin,
		"SQRT":        lvm.SquareRoot,
		"EXP":         lvm.Exp,
		"HEADING":     lvm.GetHeading,
		"DISTANCE":    lvm.GetDistance,
		"NOW":         lvm.GetTick,
		"PALLETTE":    lvm.renderPallette,
		"ATAN":        lvm.Atan,
		"PI":          lvm.Pi,
		"UID":         lvm.UID,
		"PRINTAT":     lvm.PrintAt,
		"FPS":         lvm.renderFPS,
		"BTN":         lvm.Button,
		"BG":          lvm.ShiftLayer,
	}
	for key, fn := range globals {
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

func (l LuaVM) PrintAt(L *lua.LState) int {
	l.fp.PrintAt(L.ToString(1), int(L.ToNumber(2)), int(L.ToNumber(3)))
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
	color := uint8(L.ToNumber(5))
	l.gp.Rect(math.MakeRect(x, y, w, h), color)
	return 0
}

func (l LuaVM) MakeTriangle(L *lua.LState) int {
	x0 := float64(L.ToNumber(1))
	y0 := float64(L.ToNumber(2))
	x1 := float64(L.ToNumber(3))
	y1 := float64(L.ToNumber(4))
	x2 := float64(L.ToNumber(5))
	y2 := float64(L.ToNumber(6))
	color := uint8(L.ToNumber(7))

	l.gp.Triangle(math.MakeVector(x0, y0), math.MakeVector(x1, y1), math.MakeVector(x2, y2), color)
	return 0
}
func (l LuaVM) MakeLine(L *lua.LState) int {
	x0 := float64(L.ToNumber(1))
	y0 := float64(L.ToNumber(2))
	x1 := float64(L.ToNumber(3))
	y1 := float64(L.ToNumber(4))

	c := uint8(L.ToNumber(5))
	l.gp.Line(math.MakeVector(x0, y0), math.MakeVector(x1, y1), c)
	return 0
}
func (l LuaVM) MakePoint(L *lua.LState) int {
	x := uint8(L.ToNumber(1))
	y := uint8(L.ToNumber(2))
	c := uint8(L.ToNumber(3))

	l.gp.Point(x, y, c)
	return 0
}
func (l LuaVM) MakeCircle(L *lua.LState) int {
	if l.gp == nil {
		return 0
	}
	x := float64(L.ToNumber(1))
	y := float64(L.ToNumber(2))
	r := float64(L.ToNumber(3))
	color := uint8(L.ToNumber(4))
	l.gp.Circ(math.MakeCircle(x, y, r), color)
	return 0
}

func (LuaVM) random(state *lua.LState) int {
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
func (LuaVM) Exp(state *lua.LState) int {
	n := float64(state.ToNumber(1))

	state.Push(lua.LNumber(math.Exp(n)))
	return 1
}
func (LuaVM) SquareRoot(state *lua.LState) int {
	n := float64(state.ToNumber(1))

	state.Push(lua.LNumber(math.Sqrt(n)))
	return 1
}
func (LuaVM) Sin(state *lua.LState) int {
	n := float64(state.ToNumber(1))

	state.Push(lua.LNumber(math.Sin(n)))
	return 1
}
func (LuaVM) Atan(state *lua.LState) int {
	n := float64(state.ToNumber(1))

	state.Push(lua.LNumber(math.Atan(n)))
	return 1
}
func (LuaVM) Pi(state *lua.LState) int {
	state.Push(lua.LNumber(math.Pi))
	return 1
}
func (LuaVM) UID(state *lua.LState) int {
	id, _ := uuid.NewUUID()
	state.Push(lua.LNumber(id.ID()))
	return 1
}

func (LuaVM) GetHeading(state *lua.LState) int {
	x0 := float64(state.ToNumber(1))
	y0 := float64(state.ToNumber(2))
	x1 := float64(state.ToNumber(3))
	y1 := float64(state.ToNumber(4))
	v0 := math.MakeVector(x0, y0)
	v1 := math.MakeVector(x1, y1)
	state.Push(lua.LNumber(v0.GetHeading(v1)))
	return 1
}
func (LuaVM) GetDistance(state *lua.LState) int {
	x0 := float64(state.ToNumber(1))
	y0 := float64(state.ToNumber(2))
	x1 := float64(state.ToNumber(3))
	y1 := float64(state.ToNumber(4))
	v0 := math.MakeVector(x0, y0)
	v1 := math.MakeVector(x1, y1)
	state.Push(lua.LNumber(v0.GetDistance(v1)))
	return 1
}

func (l LuaVM) Button(state *lua.LState) int {
	button := gamepad.Button(state.ToNumber(1))
	state.Push(lua.LBool(l.controllers[0].Buttons[gamepad.Button(button)]))
	return 1
}
func (l LuaVM) ShiftLayer(state *lua.LState) int {
	l.gp.ShiftLayer(uint8(state.ToNumber(1)))
	return 0
}

func (l LuaVM) GetTick(state *lua.LState) int {
	state.Push(lua.LNumber(*l.tick))
	return 1
}

func (l LuaVM) renderPallette(state *lua.LState) int {
	for i := range config.Pallette {
		x := float64(i*config.Get().Window.Width/len(config.Pallette)) + 1
		y := float64(config.Get().Window.Height - config.Get().Window.Width/len(config.Pallette))
		w := float64(config.Get().Window.Width/len(config.Pallette)) + 1
		h := float64(config.Get().Window.Width / len(config.Pallette))
		l.gp.Rect(math.MakeRect(x, y, w, h), uint8(i))
	}
	return 0
}

func (l LuaVM) renderFPS(state *lua.LState) int {
	config.Get().FPSEnabled = true
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
func (l LuaVM) UpdateCalls(state *lua.LState) {
	if err := state.CallByParam(lua.P{
		Fn:      state.GetGlobal("UPDATE"), // name of Lua function
		NRet:    0,                         // number of returned values
		Protect: true,                      // return err or panic
	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
		panic(err)
	}

	for _, g := range l.controllers {
		g.Update()
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

func (l LuaVM) initializeTick() {
	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				rand.Seed(time.Now().UnixNano())
				tick++
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
