package glrunner

// import (
// 	"image/color"
// 	"runtime"
// 	"time"
// 	"turtle/internal/graphics/glrunner/window"

// 	"github.com/go-gl/glfw/v3.3/glfw"
// )

// type System interface {
// 	Update()
// }

// type Renderer interface {
// 	Draw()
// }

// type Game struct {
// 	window          *glfw.Window
// 	Width           int
// 	Height          int
// 	WindowTitle     string
// 	WindowScale     int
// 	BackgroundColor color.Color
// 	closed          bool
// }

// func (g *Game) Run() {
// 	// GLFW event handling must run on the main OS thread
// 	runtime.LockOSThread()

// 	defer glfw.Terminate()
// 	println("running...")
// 	g.window = window.InitGL()

// 	for !g.window.ShouldClose() {
// 		<-time.Tick(time.Second / 60)
// 		g.update()
// 	}
// }

// func (g *Game) ShouldClose() bool {
// 	return g.closed
// }

// func (g *Game) update() {
// 	g.window.SwapBuffers()
// 	glfw.PollEvents()
// }
