
local T={}
local rectSize=1
local speed=.1
local initialized=false

local function make_prey(n)
    for _=0, n do
        local sizeVariance = rand(10)
        table.insert(T, {
            x=rand((windowW()-rectSize)),
            y=rand((windowH()-rectSize)),
            w=rectSize * sizeVariance,
            h=rectSize * sizeVariance,
            dx=speed,
            dy=speed,
            color=rand(15),
        })
    end
end

local function init()
    make_prey(50)
    initialized=true
end

local function move_prey()
    for _, r in pairs(T) do
        if (r.x  >= windowW()-rectSize or r.x < 0) then
            r.dx = r.dx * -1
        end

        if (r.y >= windowH()-rectSize or r.y < 0) then
            r.dy = r.dy * -1
        end

        r.x=r.x+r.dx
        r.y=r.y+r.dy
    end
end

local function render_prey()
    for _, r in pairs(T) do
        rect(r.x,r.y,r.h,r.w, r.color)
    end
end

function Update()
    if not initialized then
        init();
    end
    move_prey();
end

function Render()
    clear();
    render_prey();
end
