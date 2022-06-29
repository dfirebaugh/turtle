package turtle

import (
	"turtle/internal/pallette"
	"turtle/pkg/gamemath"
)

func Line(o Vector, d Vector, s float64, c int) {
	gp.Line(gamemath.Vector(o), gamemath.Vector(d), s, pallette.Color(c))
}