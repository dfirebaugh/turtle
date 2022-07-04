//go:build !gl
// +build !gl

package engine

import "turtle/internal/emulator/engine/ebiten"

func New(emulator ebiten.GameConsole) *ebiten.Game {
	return ebiten.New(emulator)
}
