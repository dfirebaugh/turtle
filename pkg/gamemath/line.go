package gamemath

type Line struct {
	Origin      Vector
	Destination Vector
}

func MakeLine(origin Vector, destination Vector) Line {
	return Line{
		Origin:      origin,
		Destination: destination,
	}
}
