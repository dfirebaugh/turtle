local x=0
local y=0
local color=0

function INIT()
    x=SCREENW()/2
    y=SCREENH()/2
    -- FPS()
end
function UPDATE()
    if BTN(0) then -- up
        y=y-1
    end
    if BTN(1) then -- down
        y=y+1
    end
    if BTN(2) then -- left
        x=x-1
    end
    if BTN(3) then -- right
        x=x+1
    end
    if BTN(4) then -- Z
        color=color+1
    end
    if BTN(5) then -- X
        color=color-1
    end
end
function RENDER()
    CLR()
    RECT(0, 0, SCREENW(), SCREENH(), 13) -- background color
    PARSESPRITE("111111dd11111111f6ff6fddffffffdddcccddddd111ddddd1d1dddd00d00ddd", x, y) -- some explicit sprite
    PALLETTE()

    PRINTAT("hello, world", 3, 9, 7)
end
--startSprites
--
--endSprites
