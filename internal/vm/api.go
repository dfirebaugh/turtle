package vm

import (
	"math/rand"
	"turtle/config"

	lua "github.com/yuin/gopher-lua"
)

func getScreenHeight(L *lua.LState) int {
	L.Push(lua.LNumber(config.Get().Window.Height))
	return 1
}
func getScreenWidth(L *lua.LState) int {
	L.Push(lua.LNumber(config.Get().Window.Height))
	return 1
}

// func ClearFB(f *fb.FrameBuffer) func(L *lua.LState) int {
// 	frameBuffer := f
// 	return func(L *lua.LState) int {
// 		frameBuffer.Clear()
// 		return 0
// 	}
// }
// func MakeRect(f *fb.FrameBuffer) func(L *lua.LState) int {
// 	frameBuffer := f
// 	return func(L *lua.LState) int {
// 		println(int(L.ToNumber(1) * 1000))
// 		x := float64(L.ToNumber(1))
// 		y := float64(L.ToNumber(2))
// 		w := float64(L.ToNumber(3))
// 		h := float64(L.ToNumber(4))
// 		color := int(L.ToNumber(5))
// 		frameBuffer.Rect(x-1, y-1, w+2, h+2, 1)
// 		frameBuffer.Rect(x, y, w, h, color)
// 		return 0
// 	}
// }
// func UpdateCalls(state *lua.LState) {
// 	if err := state.CallByParam(lua.P{
// 		Fn:      state.GetGlobal("Update"), // name of Lua function
// 		NRet:    0,                         // number of returned values
// 		Protect: true,                      // return err or panic
// 	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
// 		panic(err)
// 	}
// }
func DrawCalls(state *lua.LState) {
	if err := state.CallByParam(lua.P{
		Fn:      state.GetGlobal("Render"), // name of Lua function
		NRet:    0,                         // number of returned values
		Protect: true,                      // return err or panic
	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
		panic(err)
	}
}

func random(state *lua.LState) int {
	state.Push(lua.LNumber(rand.Intn(state.ToInt(1))))
	return 1
}

func scaleFactor(state *lua.LState) int {
	state.Push(lua.LNumber(10))
	return 1
}
