# The Turtle Console Emulator
Turtle is a [fantasy console](https://en.wikipedia.org/wiki/Fantasy_video_game_console) emulator.
The intention of turtle is to make it easy to make games.

### Web Editor/Emulator
You can create/edit carts for turtle script directly in the browser.
[web emulator](https://dfirebaugh.github.io/turtle/)

### Sprite Editor
Press Tab to access the sprite editor.

Sprites will be saved to the cart as a comment.

You can also render a sprite directly with `PARSESPRITE()`.
e.g.
```lua
PARSESPRITE("8888888c8777777c888888880707707c0777777cccc22ccccc7227ccccc88ccc", x, y)
```

## Turtle Carts
There is a subset of the lua scripting language embedded in turtle.
Some example turtle scripts (aka `carts`) exist in the `./examples` dir.

you can load these carts by building the project and then running the following:
```bash
./turtle -cart ./examples/simple.lua
```

Alternatively, you can run these directly in the web browser.


You must implement 3 functions:
```lua
function INIT()
-- runs on cart load
end

function UPDATE()
-- runs every update of game loop
end

function RENDER()
-- runs for draw calls
end
```

### API avialable for carts:
```lua
-- util functions
SCREENH()
SCREENW()
NOW() -- seconds since the console started
BUTTON(n) -- returns true if button is pressed - shorthand: BTN(n)
HEADING(x0, y0, x1, y1) -- heading from one point to another
DISTANCE(x0, y0, x1, y1) -- distance between two points


-- render functions
RECTANGLE(x, y, w, h, color) -- color is an index on the pallette
CIRCLE(x, y, r, color) -- shorthand: CIR(x, y, r, color)
LINE(x, y, x0, y0, x1, y1, color)
TRIANGLE(x0, y0, x1, y1, x2, y2, color) -- shorthand: TRI(x0, y0, x1, y1, x2, y2, color)
POINT(x, y, color) -- shorthand: PT(x, y, color)
CLEAR() -- clear screen - shorthand: CLR() or CLS()
SPRITE(i) -- renders a sprite at index n of a cart's sprite memory -- shorthand: SPR(i)
UID() -- generate a unique id
FPS() -- render FPS info
PRINTAT(string, x, y, color) -- print some text to the screen
PALLETTE() -- render the pallette
```

#### Object Linking
You can copy an object's properties with the following:
```lua
destination=SHALLOWCOPY(source)
```

If you want to copy an objects methods, use the clone method.


```lua
local Base=OBJECT:clone() -- clone the base class
Base.val="hello" -- set some value
function Base:say_hello() -- declare some method
    print(self.val)
end

Enemy=Base:clone() -- clone the base

Other=Base:clone() -- clone the base
Other.val="greetings" -- set a new value

Enemy:say_hello() --> "hello"
Other:say_hello() --> "greetings"
```

#### Lua 5.1 reference manual
https://www.lua.org/manual/5.1/
> note: the math library is handy (e.g. math.pi(); math.random(); math.cos(); math.sin();)

### Controls

```lua
    if BTN(0) then -- up
    end
    if BTN(1) then -- down
    end
    if BTN(2) then -- left
    end
    if BTN(3) then -- right
    end
    if BTN(4) then -- Z
    end
    if BTN(5) then -- X
    end
```


## Chips
Turtle has `chips`.  `Chips` are really just convenience libraries that are available to the `carts` (carts are intended to represent game cartridges.)


