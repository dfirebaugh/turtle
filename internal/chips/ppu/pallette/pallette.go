package pallette

import (
	"image/color"
	"math/rand"
	"turtle/config"
)

type Color uint8

var Colors = config.Pallette

func RandomColor() int {
	return rand.Intn(len(Colors))
}

func GetColor(i uint8) color.Color {
	return Colors[i%uint8(len(Colors))]
}
