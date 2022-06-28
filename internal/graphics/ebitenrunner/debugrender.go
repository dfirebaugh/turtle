package ebitenrunner

import (
	"fmt"
	"turtle/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DebugRenderer struct{}

func (d DebugRenderer) Draw(screen *ebiten.Image) {
	if !config.Get().DebugEnabled {
		return
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
}
