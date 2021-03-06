package vram

import (
	"image/color"
	"turtle/config"
	"turtle/internal/emulator/chips/ppu/pallette"
)

type graphicBuffer [config.ScreenWidth][config.ScreenHeight][3]uint8
type bufferLabel uint8

const (
	tmpBuffer bufferLabel = iota
	activeBuffer
)

type VRAM struct {
	buffers []graphicBuffer
}

func New() *VRAM {
	tb := graphicBuffer{}
	ab := graphicBuffer{}
	return &VRAM{
		buffers: []graphicBuffer{tb, ab},
	}
}

func (v *VRAM) GetFrame() []byte {
	var frame []byte = make([]byte, 0, 4*config.ScreenHeight*config.ScreenWidth)
	for y := 0; y < config.ScreenHeight; y++ {
		for x := 0; x < config.ScreenWidth; x++ {
			frame = append(frame, v.buffers[tmpBuffer][x][y][0])
			frame = append(frame, v.buffers[tmpBuffer][x][y][1])
			frame = append(frame, v.buffers[tmpBuffer][x][y][2])
			frame = append(frame, 0xFF)
		}
	}

	return frame
}

func (v *VRAM) Fill(c uint8) {
	for row := 0; row < config.Get().Window.Height; row++ {
		for col := 0; col < config.Get().Window.Width; col++ {
			r, g, b, _ := pallette.GetColor(c).RGBA()
			v.buffers[tmpBuffer][col][row][0] = uint8(r) // red
			v.buffers[tmpBuffer][col][row][1] = uint8(g) // green
			v.buffers[tmpBuffer][col][row][2] = uint8(b) // blue
		}
	}
	v.Swap()
}

func (v *VRAM) Clear() {
	v.buffers[tmpBuffer] = graphicBuffer{}
	v.buffers[activeBuffer] = v.buffers[tmpBuffer]
	v.buffers[tmpBuffer] = graphicBuffer{}
}

func (v *VRAM) Swap() {
	v.buffers[activeBuffer] = v.buffers[tmpBuffer]
}

func (v *VRAM) Put(x, y uint16, color uint8) {
	if x <= 0 || x >= config.ScreenWidth || y <= 0 || y >= config.ScreenHeight {
		return
	}
	r, g, b, _ := pallette.GetColor(color).RGBA()
	v.buffers[tmpBuffer][x][y][0] = uint8(r)
	v.buffers[tmpBuffer][x][y][1] = uint8(g)
	v.buffers[tmpBuffer][x][y][2] = uint8(b)
}

func (v *VRAM) GetBuffer() graphicBuffer {
	return v.buffers[activeBuffer]
}

func (v *VRAM) Size() (int16, int16) {
	return config.ScreenWidth, config.ScreenHeight
}

func (v *VRAM) SetPixel(x, y int16, c color.RGBA) {
	if x <= 0 || x >= config.ScreenWidth || y <= 0 || y >= config.ScreenHeight {
		return
	}
	r, g, b, _ := c.RGBA()
	v.buffers[tmpBuffer][x][y][0] = uint8(r)
	v.buffers[tmpBuffer][x][y][1] = uint8(g)
	v.buffers[tmpBuffer][x][y][2] = uint8(b)
}

func (v *VRAM) Display() error {
	return nil
}
