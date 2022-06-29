package main

// import (
// 	"turtle/pkg/turtle"
// )

// type Cart struct{}

// type rect struct {
// 	x     float64
// 	y     float64
// 	w     float64
// 	h     float64
// 	dir   float64
// 	speed float64
// 	color int
// }

// const (
// 	speed        = 100
// 	rectSize     = 1
// 	sizeVariance = 3
// )

// var (
// 	rects    = []*rect{}
// 	lastTick = 0
// )

// func main() {
// 	turtle.Run(&Cart{})
// }

// func (Cart) Init() {
// 	makeRects(15)
// 	size := 15.0
// 	r := &rect{
// 		x:     float64(turtle.Rand(turtle.ScreenWidth)) - size,
// 		y:     float64(turtle.Rand(turtle.ScreenHeight)) - size,
// 		w:     size,
// 		h:     size,
// 		speed: .78,
// 		color: turtle.RandomColor(),
// 	}
// 	r.dir = turtle.MakeVector(r.x, r.y).GetHeading(turtle.MakeVector(float64(turtle.Rand(turtle.ScreenWidth)), float64(turtle.Rand(turtle.ScreenHeight))))
// 	rects = append(rects, r)
// }

// func (Cart) Update() {
// 	moveRects()
// 	spawner()
// 	tick()
// }

// func (Cart) Render() {
// 	turtle.Clear()
// 	renderRects()
// }

// func tick() {
// 	lastTick = turtle.Tick
// }

// func spawner() {
// 	if lastTick == 0 {
// 		return
// 	}
// 	if lastTick == turtle.Tick {
// 		return
// 	}

// 	makeRects(15)
// }

// func makeRects(n int) {
// 	for i := 0; i < n; i++ {
// 		size := float64(turtle.Rand(sizeVariance)) + 1
// 		r := &rect{
// 			x:     float64(turtle.Rand(turtle.ScreenWidth)) - size,
// 			y:     float64(turtle.Rand(turtle.ScreenHeight)) - size,
// 			w:     size,
// 			h:     size,
// 			speed: .78,
// 			color: turtle.RandomColor(),
// 		}
// 		r.dir = turtle.MakeVector(r.x, r.y).GetHeading(turtle.MakeVector(float64(turtle.Rand(turtle.ScreenWidth)), float64(turtle.Rand(turtle.ScreenHeight))))
// 		rects = append(rects, r)
// 	}
// }

// func moveRects() {
// 	for i, r := range rects {
// 		collisionBody := turtle.MakeRectangle(r.x, r.y, r.w, r.h)
// 		if !collisionBody.IsWithinScreen() {
// 			r.dir = turtle.MakeVector(float64(turtle.ScreenWidth)/2, float64(turtle.ScreenHeight)/2).GetHeading(turtle.MakeVector(r.x, r.y))
// 		}

// 		r.checkCollisions(i, collisionBody)

// 		r.x += turtle.Cos(r.dir) * r.speed
// 		r.y += turtle.Sin(r.dir) * r.speed
// 	}
// }

// func (r *rect) checkCollisions(i int, collisionBody turtle.Rectangle) {
// 	for j, other := range rects {
// 		if i == j {
// 			continue
// 		}
// 		otherCB := turtle.MakeRectangle(other.x, other.y, other.w, other.h)
// 		if collisionBody.IsAxisAlignedCollision(otherCB) {
// 			if r.x < other.x && r.x > other.x+other.w {
// 				r.x = other.x + other.w
// 			}
// 			if r.y < other.y && r.y > other.y+other.h {
// 				r.y = other.y + other.h
// 			}

// 			r.dir = turtle.MakeVector(r.x, r.y).GetHeading(turtle.MakeVector(other.x, other.y))
// 		}
// 	}
// }
// func renderRects() {
// 	for _, r := range rects {
// 		// turtle.Rect(r.x-1, r.y-1, r.w+2, r.h+2, turtle.Pallette(4))
// 		// turtle.Rect(r.x, r.y, r.w, r.h, turtle.RandomColor())
// 		turtle.Circ(r.x, r.y, r.w/2, r.color)
// 	}
// }
