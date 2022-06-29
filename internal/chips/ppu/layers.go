package ppu

type GraphicsLayer [ScreenWidth][ScreenHeight][3]uint8
type Layer uint

const (
	BackgroundLayer Layer = iota
	SpriteLayer
	WindowLayer
)

func newGraphicsLayer() GraphicsLayer {
	return [ScreenWidth][ScreenHeight][3]uint8{}
}

func (gl GraphicsLayer) GetFrame() []byte {
	var frame []byte = make([]byte, 0, 4*ScreenHeight*ScreenWidth)
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			frame = append(frame, gl[x][y][0])
			frame = append(frame, gl[x][y][1])
			frame = append(frame, gl[x][y][2])
			frame = append(frame, 0xFF)
		}
	}

	return frame
}
