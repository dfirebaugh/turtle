package font

import (
	"image/color"
	"turtle/internal/emulator/chips/ppu/pallette"
	"turtle/internal/emulator/chips/vram"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

type FontProcessingUnit struct {
	vram *vram.VRAM
}

func New(v *vram.VRAM) *FontProcessingUnit {
	return &FontProcessingUnit{
		vram: v,
	}
}

func (fp *FontProcessingUnit) PrintAt(s string, x int, y int, c uint8) {
	tinyfont.WriteLine(fp.vram, &proggy.TinySZ8pt7b, int16(x), int16(y), s, pallette.GetColor(c).(color.RGBA))
}
