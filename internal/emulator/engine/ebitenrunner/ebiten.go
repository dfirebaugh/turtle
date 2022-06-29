package ebitenrunner

import (
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type System interface {
	Update()
}

type Drawable interface {
	Draw(screen *ebiten.Image)
}

type scene struct {
	systems   []System
	drawables []Drawable
}

func New(systems []System, drawables []Drawable) *scene {
	return &scene{
		systems:   systems,
		drawables: drawables,
	}
}

func (s *scene) Update() {
	for _, sys := range s.systems {
		sys.Update()
	}
	if ebiten.IsWindowBeingClosed() {
		s.Exit()
	}
}

func (s *scene) Draw(screen *ebiten.Image) {
	for _, sys := range s.drawables {
		sys.Draw(screen)
	}
	debug := DebugRenderer{}
	debug.Draw(screen)
}

func (s *scene) Exit() {
	os.Exit(0)
}
