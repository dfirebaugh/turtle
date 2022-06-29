package shapes

import (
	"turtle/internal/pallette"
	"turtle/pkg/gamemath"
)

type Shape interface {
	Draw()
}

type Circle struct {
	Color pallette.Color
	gamemath.Circle
}

type Rect struct {
	gamemath.Rect
	Color pallette.Color
}

type Line struct {
	gamemath.Line
	Color pallette.Color
}

func (Circle) Draw() {}
func (Rect) Draw()   {}
func (Line) Draw()   {}
