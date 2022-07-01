package ebitenwrapper

import (
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"
	"turtle/config"
	"turtle/internal/emulator/engine/ebitenrunner"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type scene interface {
	Reset(interface{})
	Update()
	Draw(screen *ebiten.Image)
	Exit()
}

type Game struct {
	Width  int
	Height int
	Scene  scene
	ebiten.Game
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
}

func New() *Game {
	c := config.Get()

	return &Game{
		Scene:           ebitenrunner.New(),
		WindowTitle:     c.Title,
		WindowScale:     c.ScaleFactor,
		Width:           c.Window.Width,
		Height:          c.Window.Height,
		BackgroundColor: colornames.Skyblue,
	}
}
func (g *Game) Update() error {
	g.Scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)
	g.Scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}

func (g *Game) Reset(s interface{}) {
	g.Scene.Reset(s)
}
func (g *Game) Run() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGINT,
		syscall.SIGTERM)
	go func() {
		<-sigc
		g.Exit()
	}()

	zoom := g.getZoom()
	ebiten.SetWindowSize(g.Width*zoom, g.Height*zoom)
	ebiten.SetWindowTitle(g.WindowTitle)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Exit() {
	g.Scene.Exit()
}

func (g *Game) getZoom() int {
	zoom := g.WindowScale
	if zoom == 0 {
		return 1
	}

	return zoom
}
