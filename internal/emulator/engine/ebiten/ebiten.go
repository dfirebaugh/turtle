package ebiten

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"
	"os/signal"
	"syscall"
	"turtle/config"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type GameConsole interface {
	Update()
	Render(img *image.RGBA)
}

type Game struct {
	Width  int
	Height int
	ebiten.Game
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
	Console         GameConsole
}

var img = image.NewRGBA(image.Rect(0, 0, config.Get().Window.Width, config.Get().Window.Height))

func New(console GameConsole) *Game {
	c := config.Get()

	return &Game{
		WindowTitle:     c.Title,
		WindowScale:     c.ScaleFactor,
		Width:           c.Window.Width,
		Height:          c.Window.Height,
		BackgroundColor: colornames.Skyblue,
		Console:         console,
	}
}
func (g *Game) Update() error {
	g.Console.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(g.BackgroundColor)
	g.Console.Render(img)
	screen.DrawImage(ebiten.NewImageFromImage(img), &ebiten.DrawImageOptions{})
	if config.Get().FPSEnabled {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}

func (g *Game) Reset(s interface{}) {
	g.Console = s.(GameConsole)
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
}

func (g *Game) getZoom() int {
	zoom := g.WindowScale
	if zoom == 0 {
		return 1
	}

	return zoom
}
