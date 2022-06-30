package math

// Rect a float64 slice with 4 elements []float64{x, y, width, height}
type Rect []float64

func MakeRect(x, y, width, height float64) Rect {
	return Rect{x, y, width, height}
}

func (r Rect) IsAxisAlignedCollision(other Rect) bool {
	ax := r[0]
	ay := r[1]
	aw := r[2]
	ah := r[3]

	bx := other[0]
	by := other[1]
	bw := other[2]
	bh := other[3]

	return ax < bx+bw &&
		ax+aw > bx &&
		ay < by+bh &&
		ah+ay > by
}

// Dimensions returns the total number of dimensions
func (r Rect) Dimensions() int {
	return 4
}

// Dimension returns the value of the i-th dimension
func (r Rect) Dimension(i int) float64 {
	return r[i]
}
