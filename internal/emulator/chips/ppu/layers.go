package ppu

import "turtle/config"

type GraphicsLayer [config.ScreenWidth][config.ScreenHeight][3]uint8
type Layer uint

const (
	BackgroundLayer Layer = iota
	SpriteLayer
	WindowLayer
)

func newGraphicsLayer() GraphicsLayer {
	return [config.ScreenWidth][config.ScreenHeight][3]uint8{}
}

func (gl GraphicsLayer) GetFrame() []byte {
	var frame []byte = make([]byte, 0, 4*config.ScreenHeight*config.ScreenWidth)
	for y := 0; y < config.ScreenHeight; y++ {
		for x := 0; x < config.ScreenWidth; x++ {
			frame = append(frame, gl[x][y][0])
			frame = append(frame, gl[x][y][1])
			frame = append(frame, gl[x][y][2])
			frame = append(frame, 0xFF)
		}
	}

	return frame
}
