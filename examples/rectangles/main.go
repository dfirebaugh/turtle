package main

import "turtle/pkg/turtle"

type Cart struct{}

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

func main() {
	turtle.Run(&Cart{})
}

func (Cart) Init() {
	makeRects(2000)
}

func (Cart) Update() {
	moveRects()
}

func (Cart) Render() {
	turtle.Clear()
	renderRects()
}

func makeRects(n int) {
	for i := 0; i < n; i++ {
		size := float64(turtle.Rand(sizeVariance))
		rects = append(
			rects,
			&rect{
				x:     float64(turtle.Rand(turtle.ScreenWidth)) - size,
				y:     float64(turtle.Rand(turtle.ScreenHeight)) - size,
				dx:    float64(turtle.Rand(speed)) / 100,
				dy:    float64(turtle.Rand(speed)) / 100,
				w:     size,
				h:     size,
				color: turtle.RandomColor(),
			})
	}
}

func moveRects() {
	for _, r := range rects {
		if r.x >= float64(turtle.ScreenWidth)-r.w || r.x < 0 {
			r.dx = r.dx * -1
		}

		if r.y >= float64(turtle.ScreenHeight)-r.h || r.y < 0 {
			r.dy = r.dy * -1
		}

		r.x += r.dx
		r.y += r.dy
	}
}

func renderRects() {
	for _, r := range rects {
		turtle.Rect(r.x-1, r.y-1, r.w+2, r.h+2, turtle.Pallette(4))
		turtle.Rect(r.x, r.y, r.w, r.h, r.color)
	}
}
