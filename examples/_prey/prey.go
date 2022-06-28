package main

import (
	"turtle/internal/world"
)

type preyController struct{}

func (p preyController) Update() {
	// world.Instance.EachEntity(func(handle world.EntityHandle) {
	// 	e := world.Instance.GetEntity(handle)
	// 	if e.HasTag(tag.Enemy) {
	// 		return
	// 	}
	// 	p.moveTowardCenter(handle)
	// })
}

func (p preyController) moveTowardCenter(handle world.EntityHandle) {
	// e := world.Instance.GetEntity(handle)

	// position := e.GetPosition()

	// x := config.Get().Window.Width / config.Get().ScaleFactor / 2
	// y := config.Get().Window.Height / config.Get().ScaleFactor / 2

	// center := gamemath.MakeVector(float64(x), float64(y))
	// eVec := gamemath.MakeVector(position.X, position.Y)
	// dir := e.GetDirection()
	// dir.Angle = center.GetHeading(eVec)
	// position.X += math.Cos(dir.Angle)
	// position.Y += math.Sin(dir.Angle)
}
