
local W={}
local size=15
local speed=10

local function makeCirc(n)
    for _=0, n do
        table.insert(W, {
            x=RND(SCREENW()-size),
            y=RND(SCREENH()-size),
            r=RND(15),
            dx=speed,
            dy=speed,
            color=RND(15),
        })
    end
end

local function renderCircs()
    for _, r in pairs(W) do
        CIR(r.x, r.y, r.r, r.color);
    end
end

function INIT()
    makeCirc(10)
end

function UPDATE()

end

function RENDER()
    renderCircs()
end


