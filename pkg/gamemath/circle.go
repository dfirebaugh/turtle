package gamemath

type Circle struct {
	X float64
	Y float64
	R float64
}

func MakeCircle(x, y, r float64) Circle {
	return Circle{
		X: x,
		Y: y,
		R: r,
	}
}
