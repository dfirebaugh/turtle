
local T={}
local rectSize=1
local speed=.1

local function make_prey(n)
    for _=0, n do
        local sizeVariance = RND(10)
        table.insert(T, {
            x=RND((SCREENW()-rectSize)),
            y=RND((SCREENH()-rectSize)),
            w=rectSize * sizeVariance,
            h=rectSize * sizeVariance,
            dx=speed,
            dy=speed,
            color=RND(15),
        })
    end
end

local function move_prey()
    for _, r in pairs(T) do
        if (r.x  >= SCREENW()-rectSize or r.x < 0) then
            r.dx = r.dx * -1
        end

        if (r.y >= SCREENH()-rectSize or r.y < 0) then
            r.dy = r.dy * -1
        end

        r.x=r.x+r.dx
        r.y=r.y+r.dy
    end
end

local function render_prey()
    for _, r in pairs(T) do
        RECT(r.x,r.y,r.h,r.w, r.color)
    end
end


function INIT()
    make_prey(50)
end

function UPDATE()
    move_prey();
end

function RENDER()
    CLS();
    render_prey();
end

