package math

import (
	"errors"
)

type Point struct {
	x float64
	y float64
}

type Line struct {
	slope float64
	yint  float64
}

func CreateLine(a, b Point) Line {
	slope := (b.y - a.y) / (b.x - a.x)
	yint := a.y - slope*a.x
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
		return Point{}, errors.New("The lines do not intersect")
	}
	x := (l2.yint - l.yint) / (l.slope - l2.slope)
	y := l.EvalX(l, x)
	return Point{x, y}, nil
}
