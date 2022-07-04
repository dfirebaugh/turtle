//go:build gl
// +build gl

package gamepad

type Keyboard struct{}

func (Keyboard) IsUpPressed() bool {
	return false
}
func (Keyboard) IsDownPressed() bool {
	return false
}
func (Keyboard) IsLeftPressed() bool {
	return false
}
func (Keyboard) IsLeftJustPressed() bool {
	return false
}
func (Keyboard) IsRightPressed() bool {
	return false
}
func (Keyboard) IsRightJustPressed() bool {
	return false
}
func (Keyboard) IsPrimaryPressed() bool {
	return false
}
func (Keyboard) IsSecondaryPressed() bool {
	return false
}
