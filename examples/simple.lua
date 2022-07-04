local x=0
local y=0
local color=0

function INIT()
    x=SCREENW()/2
    y=SCREENH()/2
    FPS()
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
    RECT(x, y, 10, 10, color)
    PALLETTE()
end
