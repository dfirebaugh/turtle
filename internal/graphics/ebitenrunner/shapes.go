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
	emptySubImage = emptyImage.SubImage(image.Rect(0, 0, config.Get().Window.Width, config.Get().Window.Height)).(*ebiten.Image)
)

func (r Rect) Draw() {
	ebitenutil.DrawRect(emptySubImage, r.Rect[0], r.Rect[1], r.Rect[2], r.Rect[3], pallette.Pallette{}.Get(r.Color))
}

func (l Line) DrawWithRects() {
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
		}.Draw()
	}
}

func (l Line) DrawWithPixels() {
	o := l.Line.Origin
	d := l.Line.Destination

	x := int(o[0])
	y := int(o[1])

	heading := d.GetHeading(o)

	// direction vector
	dx := int(math.Cos(heading))
	dy := int(math.Sin(heading))

	for i := 0; float64(i) < gamemath.GetDistance(l.Line.Origin, l.Line.Destination); i++ {
		emptySubImage.Set(
			x+(i*dx),
			y+(i*dy),
			pallette.Colors[l.Color],
		)
	}
}

func (l Line) DrawAsVec() {
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
	emptySubImage.DrawTriangles(vs, is, emptySubImage, op)
}

func (l Line) Draw() {
	l.DrawWithRects()
}

func (c Circle) DrawOutlineWithPixels() {
	for i := 0; i < 360; i++ {
		x1 := c.R * math.Cos(float64(i)*math.Pi/180)
		y1 := c.R * math.Sin(float64(i)*math.Pi/180)

		emptySubImage.Set(int(c.X+x1+c.R), int(c.Y+y1+c.R), pallette.Colors[c.Color])
	}
}

func (c Circle) DrawOutlineWithRects() {
	for i := 0; i < 360; i++ {
		x1 := c.R * math.Cos(float64(i)*math.Pi/180)
		y1 := c.R * math.Sin(float64(i)*math.Pi/180)

		Rect{
			Rect:  gamemath.MakeRect(c.X+x1+c.R, c.Y+y1+c.R, 1, 1),
			Color: c.Color,
		}.Draw()
	}
}

func (c Circle) DrawOutline() {
	c.DrawOutlineWithPixels()
}

// SlowFill draws a circle O(radius)
// bigger circles won't fill entirely
func (c Circle) SlowFill() {
	for radius := c.R; radius >= 0; radius-- {
		Circle{
			Circle: gamemath.MakeCircle(c.X, c.Y, radius),
			Color:  c.Color,
		}.DrawOutline()
	}
}

// fill with mostly one rect
func (c Circle) RectFill() {
	offsetX := c.R
	offsetY := c.R

	Rect{
		Rect:  gamemath.MakeRect(c.X-offsetX, c.Y-offsetY, c.R+c.R, c.R+c.R),
		Color: c.Color,
	}.Draw()
}

func (c Circle) Fill() {
	c.RectFill()
}

func (c Circle) Draw() {
	// c.Fill(dst)
	c.DrawOutline()
}

func (sr ShapeRenderer) DrawShapes() {
	if sr.Shapes == nil {
		return
	}
	for _, s := range *sr.Shapes {
		if s, ok := s.(shapes.Circle); ok {
			Circle(s).Draw()
		}
		if s, ok := s.(shapes.Rect); ok {
			Rect(s).Draw()
		}
		if s, ok := s.(shapes.Line); ok {
			Line(s).Draw()
		}
	}
}

func (sr ShapeRenderer) clear() {
	*sr.Shapes = nil
	*sr.Debug = ""
}

func (sr ShapeRenderer) debugPrint() {
	if config.Get().DebugEnabled {
		ebitenutil.DebugPrintAt(emptySubImage, fmt.Sprintf("#%d", len(*sr.Shapes)), 0, config.Get().Window.Height-15)
	}
	if *sr.Debug == "" {
		return
	}
	ebitenutil.DebugPrintAt(emptySubImage, *sr.Debug, 0, 20)
}

func (sr ShapeRenderer) Draw(dst *ebiten.Image) {
	sr.DrawShapes()
	sr.debugPrint()
	dst.DrawImage(emptySubImage, &ebiten.DrawImageOptions{})

	sr.clear()

	emptySubImage.Clear()
}
