//go:build gl
// +build gl

package engine

import (
	"turtle/config"
	"turtle/internal/emulator/engine/gl"

	"golang.org/x/image/colornames"
)

func New(emulator gl.GameConsole) *gl.Game {
	return &gl.Game{
		Width:           config.Get().Window.Width,
		Height:          config.Get().Window.Width,
		WindowTitle:     config.Get().Title,
		WindowScale:     config.Get().ScaleFactor,
		BackgroundColor: colornames.Skyblue,
		Console:         emulator,
	}
}
