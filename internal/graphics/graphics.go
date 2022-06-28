package graphics

import (
	"image/color"
	"turtle/internal/graphics/glrunner"
	"turtle/internal/graphics/pixelglrunner"
	"turtle/internal/graphics/shapes"
	"turtle/internal/pallette"
	"turtle/pkg/gamemath"
)

type scene interface {
	Update()
	Draw()
	Exit()
}

type Cart interface {
	Update()
	Render()
	Init()
}

type runner interface {
	Run()
}

type GraphicsPipeline struct {
	Width           int
	Height          int
	Scene           scene
	WindowTitle     string
	WindowScale     int
	BackgroundColor color.Color
	debug           string

	// shape buffer get cleared on ever frame
	shapes []shapes.Shape
}

func (g *GraphicsPipeline) Run(cart Cart) {
	// if config.Get().UseOpenGL {
	// 	runner := g.gl(cart)
	// 	runner.Run()
	// 	return
	// }
	// runner := g.ebiten(cart)
	// runner.Run()

	g.pixelgl(cart).Run()
}

func (g *GraphicsPipeline) Circ(circle gamemath.Circle, color pallette.Color) {
	g.shapes = append(g.shapes, shapes.Circle{
		Circle: circle,
		Color:  color,
	})
}

func (g *GraphicsPipeline) Rect(rect gamemath.Rect, color pallette.Color) {
	g.shapes = append(g.shapes, shapes.Rect{
		Rect:  rect,
		Color: color,
	})
}

func (g *GraphicsPipeline) Line(o, d gamemath.Vector, stroke float64, color pallette.Color) {
	g.shapes = append(g.shapes, shapes.Line{
		Line:  gamemath.MakeLine(o, d),
		Color: color,
	})
}

func (g *GraphicsPipeline) Print(s string) {
	g.debug = s
}

// Clear the shapebuffers when a app requests
func (g *GraphicsPipeline) Clear() {
	g.shapes = []shapes.Shape{}
}

// func (g *GraphicsPipeline) ebiten(cart Cart) runner {
// 	cartRunner := &ebitenrunner.CartRunner{Cart: cart}
// 	systems := []ebitenrunner.System{
// 		cartRunner,
// 	}

// 	drawables := []ebitenrunner.Drawable{
// 		cartRunner,
// 		ebitenrunner.ShapeRenderer{
// 			Shapes: &g.shapes,
// 			Debug:  &g.debug,
// 		},
// 	}

// 	return &ebitenwrapper.Game{
// 		Scene:           ebitenrunner.New(systems, drawables),
// 		WindowTitle:     g.WindowTitle,
// 		WindowScale:     g.WindowScale,
// 		Width:           g.Width,
// 		Height:          g.Height,
// 		BackgroundColor: g.BackgroundColor,
// 	}
// }

func (g *GraphicsPipeline) gl(cart Cart) runner {
	return &glrunner.Game{
		WindowTitle:     g.WindowTitle,
		WindowScale:     g.WindowScale,
		Width:           g.Width,
		Height:          g.Height,
		BackgroundColor: g.BackgroundColor,
	}
}
func (g *GraphicsPipeline) pixelgl(cart Cart) runner {
	cartRunner := &pixelglrunner.CartRunner{
		Cart: cart,
	}

	return &pixelglrunner.Game{
		WindowTitle:     g.WindowTitle,
		WindowScale:     g.WindowScale,
		Width:           g.Width,
		Height:          g.Height,
		BackgroundColor: g.BackgroundColor,
		Systems: []pixelglrunner.System{
			cartRunner,
		},
		Drawables: []pixelglrunner.Renderer{
			cartRunner,
			pixelglrunner.NewShapeRenderer(&g.shapes),
		},
	}
}
