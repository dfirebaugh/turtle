package cart

import (
	_ "embed"
	"turtle/config"
	"turtle/internal/cart/vm"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	lua "github.com/yuin/gopher-lua"
)

type Cart struct {
	state           *lua.LState
	gp              vm.GraphicsPipeline
	fp              vm.FontPipeline
	sm              *SpriteMemory
	lvm             vm.LuaVM
	spritelyRunning bool
	cartCode        string
	setEditorCb     func(string)
}

//go:embed spritely.lua
var spritely string

func NewCart(gp vm.GraphicsPipeline, fp vm.FontPipeline) *Cart {
	return &Cart{
		gp: gp,
		fp: fp,
	}
}

func (gr *Cart) LoadCart(cartCode string) error {
	gr.cartCode = cartCode
	gr.spritelyRunning = false
	config.Reset()
	state := lua.NewState()
	sm := &SpriteMemory{}
	sm.Sprites = make(map[uint8]string)
	sm.StoreSprites(sm.ParseSprites(gr.cartCode))

	if err := state.DoString(sm.LoadSpritesFromCart(cartCode)); err != nil {
		return err
	}
	gr.sm = sm
	gr.lvm = vm.NewLuaVM(gr.gp, gr.fp, sm)
	gr.lvm.LoadCart(state)
	gr.state = state
	gr.Init()
	return nil
}

func (gr *Cart) LoadCartFromFile(cartPath string) error {
	// TODO: need to os.Open then run gr.LoadCart

	return nil
}

func (gr *Cart) LoadSpritely() {
	gr.spritelyRunning = true
	config.Reset()
	state := lua.NewState()

	gr.sm.Sprites = make(map[uint8]string)
	gr.sm.StoreSprites(gr.sm.ParseSprites(gr.cartCode))

	if err := state.DoString(spritely); err != nil {
		println(err)
	}
	gr.lvm.LoadCart(state)
	gr.state = state

	gr.Init()
}

func (gr *Cart) SetEditorCb(fn func(string)) {
	gr.setEditorCb = fn
}

func (gr Cart) Init() {
	if gr.state == nil {
		return
	}
	gr.lvm.Init(gr.state)
}

func (gr *Cart) Update() {
	if gr.state == nil {
		return
	}
	gr.lvm.UpdateCalls(gr.state)
	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		if gr.spritelyRunning {
			gr.spritelyRunning = false
			gr.setEditorCb(gr.sm.SaveSpritesToCart(gr.cartCode))
			return
		}
		gr.LoadSpritely()
	}
}

func (gr Cart) Render() {
	if gr.state == nil {
		return
	}
	gr.lvm.DrawCalls(gr.state)
}
