package prey

import (
	"math/rand"
	"turtle/config"
	"turtle/internal/govm"
	"turtle/internal/pallette"
)

type Cart struct{}

var api *govm.GOCart

type rect struct {
	x     float64
	y     float64
	w     float64
	h     float64
	dx    float64
	dy    float64
	color int
}

const (
	speed        = 100
	rectSize     = 1
	sizeVariance = 10
)

var (
	rects = []*rect{}
)

func (Cart) Init(gc *govm.GOCart) {
	api = gc
	makePrey(100)
}

func (Cart) Update() {
	movePrey()
}

func (Cart) Render() {
	api.Clear()
	renderPrey()
}

func makePrey(n int) {
	for i := 0; i < n; i++ {
		size := float64(rand.Intn(sizeVariance))
		rects = append(
			rects,
			&rect{
				x:     float64(rand.Intn(config.Get().Window.Width)) - size,
				y:     float64(rand.Intn(config.Get().Window.Height)) - size,
				dx:    float64(rand.Intn(speed)) / 1000,
				dy:    float64(rand.Intn(speed)) / 1000,
				w:     size,
				h:     size,
				color: pallette.Pallette{}.RandomColor(),
			})
	}
}

func movePrey() {
	for _, r := range rects {
		if r.x >= float64(config.Get().Window.Width)-r.w || r.x < 0 {
			r.dx = r.dx * -1
		}

		if r.y >= float64(config.Get().Window.Height)-r.h || r.y < 0 {
			r.dy = r.dy * -1
		}

		r.x += r.dx
		r.y += r.dy
	}
}

func renderPrey() {
	for _, r := range rects {
		api.Rect(r.x, r.y, r.w, r.h, r.color)
	}
}
