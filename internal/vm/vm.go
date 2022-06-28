package vm

import (
	lua "github.com/yuin/gopher-lua"
)

var globals = map[string]lua.LGFunction{
	"screenH":     getScreenHeight,
	"screenW":     getScreenWidth,
	"rand":        random,
	"scaleFactor": scaleFactor,
}

func LoadGlobals(L *lua.LState) {
	for key, fn := range globals {
		L.SetGlobal(key, L.NewFunction(fn))
	}
}
