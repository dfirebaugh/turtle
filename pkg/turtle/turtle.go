package turtle

import (
	"image/color"
	"time"
	"turtle/config"
	"turtle/internal/graphics"
	"turtle/internal/pallette"
	"turtle/pkg/gamemath"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var (
	// Rect         = e.Rect
	RandomColor  = pallette.Pallette{}.RandomColor
	Pallette     = pallette.Pallette{}.GetColorIndex
	gp           *graphics.GraphicsPipeline
	ScreenHeight = config.Get().Window.Height
	ScreenWidth  = config.Get().Window.Width
	Tick         = 0
)

func Circ(x, y, r float64, color int) {
	gp.Circ(gamemath.MakeCircle(x, y, r), pallette.Color(color))
}

func Rect(x, y, w, h float64, color int) {
	gp.Rect(gamemath.MakeRect(x, y, w, h), pallette.Color(color))
}

func Clear() {
	gp.Clear()
}

func Print(s string) {
	gp.Print(s)
}

func MakeUID() uint32 {
	return uuid.New().ID()
}

func initializeTick() {
	ticker := time.NewTicker(time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				Tick++
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func Run(cart graphics.Cart) {
	logrus.SetLevel(logrus.ErrorLevel)
	initializeTick()

	c := config.Get()

	gp = &graphics.GraphicsPipeline{
		WindowTitle: c.Title,
		WindowScale: c.ScaleFactor,
		Width:       c.Window.Width,
		Height:      c.Window.Height,
		// BackgroundColor: e.GetColor(1),
		BackgroundColor: color.Black,
	}

	gp.Run(cart)
}
