
local a = {
    x= 0, 
    y= 0,
    r= 0,
    c= 0,
    dir=1
}

function INIT()
    a.x = 64
    a.y = 64
    a.r = 15
    a.c = 10
end

function UPDATE()
    if a.x > SCREENW()-(a.r*2) then
        a.dir = a.dir * -1
    end

    if a.x < 0 then
        a.dir = a.dir *-1
    end
    a.x= a.x +a.dir
end

function RENDER()
    clear()
    rect(64,64,10,10,1)
    line(0,0,SCREENW(), SCREENH(), 10)
    circle(64, 64, 10, 10)

    circle(a.x, a.y, a.r, a.c)
end
