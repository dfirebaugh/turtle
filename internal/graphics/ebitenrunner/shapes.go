package ebitenrunner

import (
	"fmt"
	"image"
	"math"
	"turtle/config"
	"turtle/internal/graphics/shapes"
	"turtle/internal/pallette"
	"turtle/pkg/gamemath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Shape interface {
	Draw(dst *ebiten.Image)
}

type Circle struct {
	Color pallette.Color
	gamemath.Circle
}

type Rect struct {
	gamemath.Rect
	Color pallette.Color
}

type Line struct {
	gamemath.Line
	Color pallette.Color
}

type ShapeRenderer struct {
	Shapes *[]shapes.Shape
	Debug  *string
}

var (
	emptyImage = ebiten.NewImage(config.Get().Window.Width, config.Get().Window.Height)

	// emptySubImage is an internal sub image of emptyImage.
	// Use emptySubImage at DrawTriangles instead of emptyImage in order to avoid bleeding edges.
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, config.Get().Window.Width, config.Get().Window.Height)).(*ebiten.Image)
)

func (r Rect) Draw(dst *ebiten.Image) {
	ebitenutil.DrawRect(dst, r.Rect[0], r.Rect[1], r.Rect[2], r.Rect[3], pallette.Pallette{}.Get(r.Color))
}

func (l Line) DrawWithRects(dst *ebiten.Image) {
	o := l.Line.Origin
	d := l.Line.Destination

	x := o[0]
	y := o[1]

	heading := d.GetHeading(o)

	dx := math.Cos(heading)
	dy := math.Sin(heading)

	for i := 0; float64(i) < gamemath.GetDistance(l.Line.Origin, l.Line.Destination); i++ {
		Rect{
			Rect:  gamemath.MakeRect(x+(float64(i)*dx), y+(float64(i)*dy), 1, 1),
			Color: l.Color,
		}.Draw(dst)
	}
}

func (l Line) DrawWithPixels(dst *ebiten.Image) {
	o := l.Line.Origin
	d := l.Line.Destination

	x := int(o[0])
	y := int(o[1])

	heading := d.GetHeading(o)

	// direction vector
	dx := int(math.Cos(heading))
	dy := int(math.Sin(heading))

	for i := 0; float64(i) < gamemath.GetDistance(l.Line.Origin, l.Line.Destination); i++ {
		dst.Set(
			x+(i*dx),
			y+(i*dy),
			pallette.Colors[l.Color],
		)
	}
}

func (l Line) DrawAsVec(dst *ebiten.Image) {
	var path vector.Path

	path.MoveTo(1, 12)
	path.LineTo(1, 1)
	path.LineTo(64, 64)
	path.LineTo(120, 120)

	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}
	r, g, b, _ := pallette.Colors[l.Color].RGBA()

	// if there are less than 3 elements, nothing is returned
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = float32(r)
		vs[i].ColorG = float32(g)
		vs[i].ColorB = float32(b)
	}
	dst.DrawTriangles(vs, is, emptySubImage, op)
}

func (l Line) Draw(dst *ebiten.Image) {
	l.DrawWithRects(dst)
}

func (c Circle) DrawOutlineWithPixels(dst *ebiten.Image) {
	for i := 0; i < 360; i++ {
		x1 := c.R * math.Cos(float64(i)*math.Pi/180)
		y1 := c.R * math.Sin(float64(i)*math.Pi/180)

		dst.Set(int(c.X+x1+c.R), int(c.Y+y1+c.R), pallette.Colors[c.Color])
	}
}

func (c Circle) DrawOutlineWithRects(dst *ebiten.Image) {
	for i := 0; i < 360; i++ {
		x1 := c.R * math.Cos(float64(i)*math.Pi/180)
		y1 := c.R * math.Sin(float64(i)*math.Pi/180)

		Rect{
			Rect:  gamemath.MakeRect(c.X+x1+c.R, c.Y+y1+c.R, 1, 1),
			Color: c.Color,
		}.Draw(dst)
	}
}

func (c Circle) DrawOutline(dst *ebiten.Image) {
	c.DrawOutlineWithPixels(dst)
}

// SlowFill draws a circle O(radius)
// bigger circles won't fill entirely
func (c Circle) SlowFill(dst *ebiten.Image) {
	for radius := c.R; radius >= 0; radius-- {
		Circle{
			Circle: gamemath.MakeCircle(c.X, c.Y, radius),
			Color:  c.Color,
		}.DrawOutline(dst)
	}
}

// fill with mostly one rect
func (c Circle) RectFill(dst *ebiten.Image) {
	offsetX := c.R
	offsetY := c.R

	Rect{
		Rect:  gamemath.MakeRect(c.X-offsetX, c.Y-offsetY, c.R+c.R, c.R+c.R),
		Color: c.Color,
	}.Draw(dst)
}

func (c Circle) Fill(dst *ebiten.Image) {
	c.RectFill(dst)
}

func (c Circle) Draw(dst *ebiten.Image) {
	// c.Fill(dst)
	c.DrawOutline(dst)
}

func (sr ShapeRenderer) DrawShapes(dst *ebiten.Image) {
	if sr.Shapes == nil {
		return
	}
	for _, s := range *sr.Shapes {
		s, ok := s.(Shape)
		if !ok {
			continue
		}
		s.Draw(dst)
	}
}

func (sr ShapeRenderer) clear() {
	*sr.Shapes = nil
	*sr.Debug = ""
}

func (sr ShapeRenderer) debugPrint(dst *ebiten.Image) {
	if *sr.Debug == "" {
		return
	}
	ebitenutil.DebugPrintAt(dst, *sr.Debug, 0, 20)
}

func (sr ShapeRenderer) Draw(dst *ebiten.Image) {
	sr.DrawShapes(dst)
	sr.debugPrint(dst)
	if config.Get().DebugEnabled {
		ebitenutil.DebugPrintAt(dst, fmt.Sprintf("#%d", len(*sr.Shapes)), 0, config.Get().Window.Height-15)
	}
	sr.clear()
}
