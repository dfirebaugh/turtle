package pixelglrunner

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type System interface {
	Update()
}

type Renderer interface {
	Draw(*pixelgl.Window)
}

type Game struct {
	Width           int
	Height          int
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
	closed          bool
	Systems         []System
	Drawables       []Renderer
	window          *pixelgl.Window
}

func (g *Game) Run() {
	pixelgl.Run(g.run)
}

func (g *Game) ShouldClose() bool {
	return g.closed
}

func (g *Game) run() {
	cfg := pixelgl.WindowConfig{
		Title:     "turtle engine",
		Bounds:    pixel.R(0, 0, float64(g.Width), float64(g.Height)),
		Resizable: true,
		VSync:     true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	g.window = win

	for !win.Closed() {
		g.Render()
		g.window.Update()
		g.Update()
	}
}

func (g *Game) Update() {
	// <-time.Tick(time.Second / 60)
	for _, r := range g.Systems {
		r.Update()
	}
}
func (g *Game) Render() {
	g.window.Clear(colornames.Black)

	for _, r := range g.Drawables {
		r.Draw(g.window)
	}
}
