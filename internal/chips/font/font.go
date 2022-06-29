package font

type FontProcessingUnit struct{}

func (fp *FontProcessingUnit) PrintAt(s string, x int, y int) {
	println(s, x, y)
}
