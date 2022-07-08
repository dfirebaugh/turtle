package cart

import (
	_ "embed"
	"os"
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
	cartPath        string
	setEditorCb     func(string)
}

//go:embed spritely.lua
var spritely string

// some lua code that is loaded into global scope
//go:embed object_linkage.lua
var objectLinkage string

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

	if err := state.DoString(objectLinkage); err != nil {
		return err
	}
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
	gr.cartPath = cartPath
	cart, err := os.ReadFile(cartPath)
	if err != nil {
		return err
	}

	gr.LoadCart(string(cart))
	return nil
}

func (gr *Cart) LoadSpritely() {
	gr.spritelyRunning = true
	config.Reset()
	state := lua.NewState()

	gr.sm.Sprites = make(map[uint8]string)
	gr.sm.StoreSprites(gr.sm.ParseSprites(gr.cartCode))
	if err := state.DoString(objectLinkage); err != nil {
		println(err)
	}
	if err := state.DoString(spritely); err != nil {
		println(err)
	}
	gr.lvm.LoadCart(state)
	gr.state = state

	gr.Init()
}

func (gr *Cart) WriteCartFile() {
	// write the sprite memory to the cart file
	f, err := os.OpenFile(gr.cartPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 07644)
	if err != nil {
		println(err.Error(), gr.cartPath)
		return
	}
	defer f.Close()

	f.Truncate(0)
	_, err = f.Seek(0, 0)
	if err != nil {
		println(err.Error())
		return
	}

	_, err = f.Write([]byte(gr.sm.SaveSpritesToCart(gr.cartCode)))
	if err != nil {
		println(err.Error())
	}
	gr.cartCode = gr.sm.SaveSpritesToCart(gr.cartCode)
	println(gr.cartCode)
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
			if gr.setEditorCb == nil {
				gr.WriteCartFile()
				gr.LoadCart(gr.cartCode)
				return
			}
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
