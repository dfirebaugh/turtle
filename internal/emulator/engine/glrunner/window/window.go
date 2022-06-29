package window

import (
	"turtle/config"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func InitGL() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	// glfw.WindowHint(glfw.Resizable, glfw.False)
	// glfw.WindowHint(glfw.ContextVersionMajor, 4)
	// glfw.WindowHint(glfw.ContextVersionMinor, 1)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(config.Get().Window.Width, config.Get().Window.Height, "Conway's Game of Life", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

// func setExclusiveMouse(exclusive bool) {
// 	if exclusive {
// 		Win.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
// 	} else {
// 		Win.SetInputMode(glfw.CursorMode, glfw.CursorNormal)
// 	}
// 	// input.ExclusiveMouse = exclusive
// }

// func onFrameBufferSizeCallback(window *glfw.Window, width, height int) {
// 	// gl.Viewport(0, 0, int32(width), int32(height))
// }
