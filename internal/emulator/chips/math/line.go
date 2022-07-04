package math

import (
	"errors"
)

type Point struct {
	X float64
	Y float64
}

type Line struct {
	slope float64
	yint  float64
}

func CreatePoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func CreateLine(a, b Point) Line {
	slope := (b.Y - a.Y) / (b.X - a.X)
	yint := a.Y - slope*a.X
	return Line{slope, yint}
}

func (Line) EvalX(l Line, x float64) float64 {
	return l.slope*x + l.yint
}

func (l Line) IsParrallel(l1, l2 Line) bool {
	return l1.slope == l2.slope
}

func (l Line) Intersection(l2 Line) (Point, error) {
	if l.slope == l2.slope {
		return Point{}, errors.New("the lines do not intersect")
	}
	x := (l2.yint - l.yint) / (l.slope - l2.slope)
	y := l.EvalX(l, x)

	println("i", int(x), int(y))
	return Point{x, y}, nil
}
