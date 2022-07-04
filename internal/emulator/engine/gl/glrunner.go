package gl

import (
	"image"
	"image/color"
	"runtime"
	"time"
	"turtle/config"
	"turtle/internal/emulator/engine/gl/window"
	"turtle/pkg/turtle"

	"github.com/disintegration/imaging"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"golang.org/x/image/draw"
)

type GameConsole interface {
	Update()
	Render(img *image.RGBA)
}

type System interface {
	Update()
}

type Renderer interface {
	Draw()
}

type Game struct {
	Width           int
	Height          int
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
	Console         GameConsole
}

var frames = 0.0
var last = 0.0

func (g *Game) Run() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()

	win := window.InitGLFW()
	defer glfw.Terminate()

	if win == nil {
		return
	}

	window.InitGL()
	window.InitFrameBuffer()

	for !win.ShouldClose() {
		g.render(win)
		go g.calculateFPS()
		go g.Console.Update()
		<-time.Tick(time.Second / 120)
	}
}

func (g *Game) Reset(interface{}) {

}

var img = image.NewRGBA(image.Rect(0, 0, config.Get().Window.Width, config.Get().Window.Height))
var scaledImg = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X*config.Get().ScaleFactor, img.Bounds().Max.Y*config.Get().ScaleFactor))

func (g *Game) render(win *glfw.Window) {
	c := config.Get()
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	// gl.UseProgram(prog)
	w, h := win.GetSize()
	g.Console.Render(img)
	gl.BindTexture(gl.TEXTURE_2D, window.Texture)

	// Set the expected size that you want:
	// Resize:
	draw.NearestNeighbor.Scale(scaledImg, scaledImg.Rect, img, img.Bounds(), draw.Over, nil)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA8, int32(c.Window.Width*c.ScaleFactor), int32(c.Window.Height*c.ScaleFactor), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(imaging.FlipV(scaledImg).Pix))
	gl.BlitFramebuffer(0, 0, int32(w), int32(h), 0, 0, int32(c.Window.Width*c.ScaleFactor), int32(c.Window.Height*c.ScaleFactor), gl.COLOR_BUFFER_BIT, gl.LINEAR)

	glfw.PollEvents()
	win.SwapBuffers()
}

func (g *Game) calculateFPS() {
	current := glfw.GetTime()
	frames++
	if current-last >= 1.0 {
		turtle.SetFPS(frames)
		frames = 0
		last += 1.0
	}
}
