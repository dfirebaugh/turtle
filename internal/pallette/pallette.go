package pallette

import (
	"image/color"
	"math/rand"
)

type Color int

const ()

type Pallette struct{}

var Colors = []color.Color{
	color.RGBA{127, 36, 84, 255},
	color.RGBA{28, 43, 83, 255},
	color.RGBA{0, 135, 81, 255},
	color.RGBA{171, 82, 54, 255},
	color.RGBA{96, 88, 79, 255},
	color.RGBA{195, 195, 198, 255},
	color.RGBA{255, 241, 233, 255},
	color.RGBA{237, 27, 81, 255},
	color.RGBA{250, 162, 27, 255},
	color.RGBA{247, 236, 47, 255},
	color.RGBA{93, 187, 77, 255},
	color.RGBA{81, 166, 220, 255},
	color.RGBA{131, 118, 156, 255},
	color.RGBA{241, 118, 166, 255},
	color.RGBA{252, 204, 171, 255},
}

func (Pallette) RandomColor() int {
	return rand.Intn(len(Colors))
}

func (p Pallette) GetColorIndex(i int) int {
	if i > len(Colors) {
		return 0
	}
	return i
}
func (Pallette) GetColor(i int) color.Color {
	return Colors[i]
}
func (Pallette) Get(i Color) color.Color {
	return Colors[i]
}