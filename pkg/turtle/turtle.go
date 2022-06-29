package turtle

import (
	"turtle/config"
	"turtle/internal/pallette"
)

var (
	RandomColor  = pallette.Pallette{}.RandomColor
	Pallette     = pallette.Pallette{}.GetColorIndex
	ScreenHeight = config.Get().Window.Height
	ScreenWidth  = config.Get().Window.Width
	Tick         = 0
)

// func Circ(x, y, r float64, color int) {
// 	gp.Circ(gamemath.MakeCircle(x, y, r), pallette.Color(color))
// }

// func Rect(x, y, w, h float64, color int) {
// 	gp.Rect(gamemath.MakeRect(x, y, w, h), pallette.Color(color))
// }

// func Clear() {
// 	gp.Clear()
// }

// func Print(s string) {
// 	gp.Print(s)
// }

// func MakeUID() uint32 {
// 	return uuid.New().ID()
// }

// func initializeTick() {
// 	ticker := time.NewTicker(time.Second)
// 	quit := make(chan struct{})
// 	go func() {
// 		for {
// 			select {
// 			case <-ticker.C:
// 				Tick++
// 			case <-quit:
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()
// }

// func RunCart(cartPath string) {
// 	Run(cart.NewCart(cartPath, gp))
// }

// func Run(cart graphics.Cart) {
// 	logrus.SetLevel(logrus.ErrorLevel)
// 	initializeTick()
// 	if gp == nil {
// 		initGP()
// 	}
// 	gp.Run(cart)
// }
