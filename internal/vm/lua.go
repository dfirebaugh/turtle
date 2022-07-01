package vm

import (
	"math"
	"math/rand"
	"time"
	"turtle/config"
	"turtle/internal/pallette"
	"turtle/pkg/gamemath"

	"github.com/google/uuid"
	lua "github.com/yuin/gopher-lua"
)

type GraphicsPipeline interface {
	Rect(rect gamemath.Rect, color pallette.Color)
	Line(v0 gamemath.Vector, v1 gamemath.Vector, stroke float64, color pallette.Color)
	Circ(circle gamemath.Circle, color pallette.Color)
	Point(x uint8, y uint8, color pallette.Color)
	Clear()
}

type FontPipeline interface {
	PrintAt(string, int, int)
}

type LuaVM struct {
	gp      GraphicsPipeline
	fp      FontPipeline
	globals map[string]lua.LGFunction
	tick    *uint
}

// world clock
var tick uint = 1

func NewLuaVM(gp GraphicsPipeline, fp FontPipeline) LuaVM {
	lvm := LuaVM{
		gp:   gp,
		fp:   fp,
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
		"POINT":       lvm.MakePoint,
		"CLS":         lvm.Clear,
		"CLR":         lvm.Clear,
		"COS":         lvm.Cos,
		"SIN":         lvm.Sin,
		"HEADING":     lvm.GetHeading,
		"DISTANCE":    lvm.GetDistance,
		"NOW":         lvm.GetTick,
		"PALLETTE":    lvm.renderPallette,
		"ATAN":        lvm.Atan,
		"PI":          lvm.Pi,
		"UID":         lvm.UID,
		"PRINTAT":     lvm.PrintAt,
		"FPS":         lvm.renderFPS,
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
	color := int(L.ToNumber(5))
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
func (l LuaVM) MakePoint(L *lua.LState) int {
	x := uint8(L.ToNumber(1))
	y := uint8(L.ToNumber(2))
	c := pallette.Color(L.ToNumber(3))

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
	color := int(L.ToNumber(4))
	l.gp.Circ(gamemath.MakeCircle(x, y, r), pallette.Color(color))
	return 0
}

func (LuaVM) random(state *lua.LState) int {
	// rand.Seed(time.Now().UnixNano())
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
	v0 := gamemath.MakeVector(x0, y0)
	v1 := gamemath.MakeVector(x1, y1)
	state.Push(lua.LNumber(v0.GetHeading(v1)))
	return 1
}
func (LuaVM) GetDistance(state *lua.LState) int {
	x0 := float64(state.ToNumber(1))
	y0 := float64(state.ToNumber(2))
	x1 := float64(state.ToNumber(3))
	y1 := float64(state.ToNumber(4))
	v0 := gamemath.MakeVector(x0, y0)
	v1 := gamemath.MakeVector(x1, y1)
	state.Push(lua.LNumber(v0.GetDistance(v1)))
	return 1
}

func (l LuaVM) GetTick(state *lua.LState) int {
	state.Push(lua.LNumber(*l.tick))
	return 1
}

func (l LuaVM) renderPallette(state *lua.LState) int {
	w := float64(config.Get().Window.Width / len(pallette.Colors))
	for i := range pallette.Colors {
		x := float64(i)*w + w
		y := float64(config.Get().Window.Height) - w
		h := w
		l.gp.Rect(gamemath.MakeRect(x, y, w, h), pallette.Color(i))
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
