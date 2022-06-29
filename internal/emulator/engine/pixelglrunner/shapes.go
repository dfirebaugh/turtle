package pixelglrunner

// import (
// 	"turtle/internal/graphics/shapes"
// 	"turtle/internal/pallette"

// 	"github.com/faiface/pixel"
// 	"github.com/faiface/pixel/imdraw"
// 	"github.com/faiface/pixel/pixelgl"
// )

// type shape interface {
// 	Draw()
// }
// type ShapeRenderer struct {
// 	shapes *[]shapes.Shape
// 	imd    *imdraw.IMDraw
// }

// func NewShapeRenderer(s *[]shapes.Shape) ShapeRenderer {
// 	imd := imdraw.New(nil)

// 	return ShapeRenderer{
// 		imd:    imd,
// 		shapes: s,
// 	}
// }

// func DrawCircle(c shapes.Circle) {

// }

// func (s ShapeRenderer) DrawCircle(c shapes.Circle) {
// 	s.imd.Color = pallette.Colors[c.Color]
// 	s.imd.Push(pixel.V(c.X+c.R, c.Y+c.R))
// 	s.imd.Circle(c.R, 0)
// }
// func (s ShapeRenderer) DrawLine(l shapes.Line) {
// 	s.imd.Color = pallette.Colors[l.Color]
// 	s.imd.Push(pixel.V(l.Origin[0], l.Origin[1]), pixel.V(l.Destination[0], l.Destination[1]))
// 	s.imd.EndShape = imdraw.RoundEndShape
// 	s.imd.Line(1)
// }
// func (s ShapeRenderer) DrawRect(r shapes.Rect) {
// 	s.imd.Color = pallette.Colors[r.Color]
// 	s.imd.Push(pixel.V(r.Rect[0], r.Rect[1]), pixel.V(r.Rect[0]+r.Rect[2], r.Rect[1]+r.Rect[3]))
// 	s.imd.Rectangle(0)
// }

// func (s ShapeRenderer) debugImg() {
// 	s.imd.Color = pixel.RGB(1, 0, 0)
// 	s.imd.Push(pixel.V(20, 10))
// 	s.imd.Color = pixel.RGB(0, 1, 0)
// 	s.imd.Push(pixel.V(80, 10))
// 	s.imd.Color = pixel.RGB(0, 0, 1)
// 	s.imd.Push(pixel.V(50, 70))
// 	s.imd.Polygon(0)
// }

// func (s ShapeRenderer) drawShapes() {
// 	for _, shp := range *s.shapes {
// 		if c, ok := shp.(shapes.Circle); ok {
// 			s.DrawCircle(c)
// 		}
// 		if l, ok := shp.(shapes.Line); ok {
// 			s.DrawLine(l)
// 		}
// 		if r, ok := shp.(shapes.Rect); ok {
// 			s.DrawRect(r)
// 		}
// 	}
// }

// func (s ShapeRenderer) clear() {
// 	*s.shapes = nil
// }

// func (s ShapeRenderer) Draw(win *pixelgl.Window) {
// 	// s.debugImg()
// 	s.imd.Clear()
// 	s.drawShapes()
// 	s.imd.Draw(win)
// 	s.clear()
// }
