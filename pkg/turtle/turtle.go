package turtle

import "turtle/config"

var (
	ScreenHeight = config.Get().Window.Height
	ScreenWidth  = config.Get().Window.Width
	Tick         = 0
)
var fps float64

func SetFPS(f float64) {
	fps = f
}

func CurrentFPS() int {
	return int(fps)
}
