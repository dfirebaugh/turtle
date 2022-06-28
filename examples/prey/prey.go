package main

import (
	"turtle/pkg/turtle"
)

type Cart struct{}

type predator struct {
	x     float64
	y     float64
	w     float64
	h     float64
	color int
	sight float64
	dir   float64
	speed float64
	ate   int
	alive bool
	life  int
	key   uint32
}

type prey struct {
	x     float64
	y     float64
	w     float64
	h     float64
	color int
	sight float64
	dir   float64
	speed float64
	alive bool
	key   uint32
}

const (
	numPredators   = 1
	numPrey        = 50
	spawnRate      = 500
	deathRate      = 50
	speedReduction = .048
)

var (
	sizeVariance = (turtle.ScreenWidth + (turtle.ScreenWidth / 2)) / (turtle.ScreenWidth / 2)
	preys        = map[uint32]*prey{}
	predators    = map[uint32]*predator{}
	tick         = 0
)

func main() {
	turtle.Run(&Cart{})
}

func (Cart) Init() {
	makePrey(numPrey)
	makePredators(numPredators)
}

func (Cart) Update() {
	updatePrey()
	updatePredators()

	tick++
}

func (Cart) Render() {
	turtle.Clear()
	renderPrey()
	renderPredator()
}

func makePreyAt(n int, x float64, y float64) {
	for i := 0; i < n; i++ {
		size := float64(turtle.Rand(sizeVariance)) + float64(sizeVariance)
		p := &prey{
			x:     x,
			y:     y,
			w:     size,
			h:     size,
			color: turtle.Pallette(2),
			sight: 10,
			alive: true,
			speed: .1,
			key:   turtle.MakeUID(),
		}
		p.dir = turtle.MakeVector(p.x, p.y).GetHeading(turtle.MakeVector(float64(turtle.Rand(turtle.ScreenWidth))-size, float64(turtle.Rand(turtle.ScreenHeight))-size))
		// preys = append(preys, p)
		preys[p.key] = p
	}
}

func makePrey(n int) {
	for i := 0; i < n; i++ {
		size := float64(turtle.Rand(sizeVariance)) + 2
		makePreyAt(1, float64(turtle.Rand(turtle.ScreenWidth))-size, float64(turtle.Rand(turtle.ScreenHeight))-size)
	}
}

func makePredatorsAt(n int, x float64, y float64) {
	for i := 0; i < n; i++ {
		p := &predator{
			x:     x,
			y:     y,
			w:     1,
			h:     1,
			color: turtle.Pallette(7),
			sight: 40,
			speed: 1,
			ate:   0,
			life:  0,
			alive: true,
			key:   turtle.MakeUID(),
		}
		p.dir = turtle.MakeVector(p.x, p.y).GetHeading(turtle.MakeVector(float64(turtle.Rand(turtle.ScreenWidth)), float64(turtle.Rand(turtle.ScreenHeight))))
		// predators = append(predators, p)
		predators[p.key] = p
	}
}

func makePredators(n int) {
	for i := 0; i < n; i++ {
		makePredatorsAt(n, float64(turtle.Rand(turtle.ScreenWidth)), float64(turtle.Rand(turtle.ScreenHeight)))
	}
}

func eachPrey(e func(p *prey)) {
	for _, r := range preys {
		e(r)
	}
}

func updatePrey() {
	eachPrey(func(r *prey) {
		if !r.alive {
			return
		}
		r.move()
		if tick%spawnRate == 0 {
			r.spawn()
		}
	})
	if tick%spawnRate == 0 {
		tick = 0
	}
}

func renderPrey() {
	eachPrey(func(r *prey) {
		r.render()
	})
}

func (r prey) render() {
	if !r.alive {
		return
	}
	// turtle.Rect(r.x, r.y, r.w, r.h, r.color)
	// turtle.Circ(r.x, r.y, r.w+1, 1)
	turtle.Circ(r.x, r.y, r.w/2, r.color)
}

func (r prey) spawn() {
	if tick == 0 {
		return
	}
	makePreyAt(1, r.x, r.y)
}

func (r *prey) moveToward(v turtle.Vector) {
	r.dir = v.GetHeading(turtle.MakeVector(r.x, r.y))
}

func (r *prey) move() {
	if !turtle.MakeRectangle(r.x, r.y, r.w, r.h).IsWithinScreen() {
		r.moveToward(turtle.MakeVector(float64(turtle.ScreenWidth)/2, float64(turtle.ScreenHeight)/2))
	}

	r.x += turtle.Cos(r.dir) * r.speed
	r.y += turtle.Sin(r.dir) * r.speed
}
func (r *prey) kill() {
	r.alive = false
	delete(preys, r.key)
}

func eachPredator(e func(p *predator)) {
	for _, r := range predators {
		e(r)
	}
}
func updatePredators() {
	eachPredator(func(r *predator) {
		if !r.alive {
			return
		}
		if r.life > turtle.ScreenWidth/4 {
			r.kill()
		}
		r.hunt()
		if tick%10 == 0 {
			r.life++
		}
	})
}

func renderPredator() {
	eachPredator(func(r *predator) {
		if !r.alive {
			return
		}
		r.render()
	})
}

func (r predator) render() {
	r.renderNose()
	turtle.Circ(r.x+r.w/2, r.y+r.h/2, r.w/2, r.color)
}

func (r predator) renderNose() {
	x := (turtle.Cos(r.dir))
	y := (turtle.Sin(r.dir))

	o := turtle.MakeVector(r.x+(r.w/2), r.y+(r.h/2))
	d := turtle.MakeVector(x*r.w, y*r.h).Add(o)

	turtle.Line(o, d, 1, turtle.Pallette(7))
}

// look for nearby prey to eat
func (r *predator) hunt() {
	collisionBody := turtle.MakeRectangle(r.x, r.y, r.w, r.h)

	found := false
	eachPrey(func(p *prey) {
		if !p.alive {
			return
		}
		if found {
			return
		}
		pPosition := turtle.Vector{p.x, p.y}
		pCollisionBody := turtle.MakeRectangle(p.x, p.y, p.w, p.h)

		if turtle.MakeRectangle(r.x, r.y, r.w+r.sight, r.h+r.sight).IsAxisAlignedCollision(pCollisionBody) {
			r.moveToward(pPosition)
			found = true
			r.move()
			if collisionBody.IsAxisAlignedCollision(pCollisionBody) {
				r.eat(p)
				return
			}

			return
		}
	})
	if found {
		return
	}
	r.move()
}

func (r *predator) moveToward(v turtle.Vector) {
	r.dir = v.GetHeading(turtle.MakeVector(r.x, r.y))
}

func (r *predator) move() {
	if !turtle.MakeRectangle(r.x, r.y, r.w, r.h).IsWithinScreen() {
		r.moveToward(turtle.MakeVector(float64(turtle.ScreenWidth)/2, float64(turtle.ScreenHeight)/2))
	}

	r.x += turtle.Cos(r.dir) * r.speed
	r.y += turtle.Sin(r.dir) * r.speed
}

func (r *predator) eat(p *prey) {
	p.kill()
	r.w += (p.w * .1)
	r.h += (p.h * .1)
	r.ate++
	r.life = 0
	if r.speed > .1 {
		r.speed -= speedReduction
	}

	if r.ate%10 == 0 {
		r.spawn()
	}
}

func (r *predator) spawn() {
	makePredatorsAt(1, r.x, r.y)
}
func (r *predator) kill() {
	r.alive = false
	delete(predators, r.key)
}
