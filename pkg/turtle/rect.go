package turtle

// Rectangle a float64 slice with 4 elements []float64{x, y, width, height}
type Rectangle []float64

func MakeRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{x, y, width, height}
}

func (r Rectangle) IsAxisAlignedCollision(other Rectangle) bool {
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
func (r Rectangle) Dimensions() int {
	return 4
}

// Dimension returns the value of the i-th dimension
func (r Rectangle) Dimension(i int) float64 {
	return r[i]
}

func (r Rectangle) Draw(color int) {

}

func (r Rectangle) IsWithinScreen() bool {
	x := r[0]
	y := r[1]
	w := r[2]
	h := r[3]
	return x+w < float64(ScreenWidth) && x > 0 && y+h < float64(ScreenHeight) && y > 0
}
