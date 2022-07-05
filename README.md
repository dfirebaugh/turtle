# The Turtle Console Emulator
Turtle is a [fantasy console](https://en.wikipedia.org/wiki/Fantasy_video_game_console) emulator.
The intention of turtle is to make it easy to make games.

### Web Editor/Emulator
You can create/edit carts for turtle script directly in the browser.
[web demo](https://dfirebaugh.github.io/turtle/)

(note: you will likely have to press reset after it loads)

## Turtle Scripting
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

api avialable in lua:
```lua
SCREENH()
SCREENW()
BUTTON(n) -- returns true if button is pressed - shorthand: BTN(n)
RECTANGLE(x, y, w, h, color) -- color is an index on the pallette
CIRCLE(x, y, r, color) -- shorthand: CIR(x, y, r, color)
LINE(x, y, x0, y0, x1, y1, color)
TRIANGLE(x0, y0, x1, y1, x2, y2, color) -- shorthand: TRI(x0, y0, x1, y1, x2, y2, color)
POINT(x, y, color) -- shorthand: PT(x, y, color)
CLEAR() -- clear screen - shorthand: CLR() or CLS()
RANDOM(n) -- random number between 0 and n - shorthand: RND()
COS(n) -- cosin
SIN(n) -- sin
SQRT(n)
ATAN(n)
PI()
HEADING(x0, y0, x1, y1) -- heading from one point to another
DISTANCE(x0, y0, x1, y1) -- distance between two points
NOW() -- seconds since the console started
UID() -- generate a unique id
FPS() -- render FPS info
PALLETTE() -- render the pallette
```

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


