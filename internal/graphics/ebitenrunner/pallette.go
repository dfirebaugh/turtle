package ebitenrunner

import (
	"turtle/internal/pallette"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var legend = ebiten.NewImage(1, 15)

func init() {
	for i, c := range pallette.Colors {
		ebitenutil.DrawRect(legend, 0, float64(i), 1, 1, c)
	}
}

func DrawPallette(dst *ebiten.Image, p pallette.Pallette) {
	dst.DrawImage(legend, &ebiten.DrawImageOptions{})
}
