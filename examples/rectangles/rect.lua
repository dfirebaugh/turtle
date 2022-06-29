local World={}
local shouldRenderPallette=true
local nRects=40

local Rect={}
function Rect:new() 
    this = {}
    this.size=RND(8)
    this.x=RND(SCREENW())
    this.y=RND(SCREENH())
    this.dir = HEADING(
        this.x,
        this.y,
        RND(SCREENW()),
        RND(SCREENH())
    )
    this.speed=0.4

    function this:render()
        RECT(this.x-1, this.y-1, this.size+2, this.size+2, RND(1))
        RECT(this.x, this.y, this.size, this.size, RND(14))
    end
    function this:isValidHeading()
        if this.dir ~= this.dir then
            return false
        end

        return true
    end
    function this:moveToward(hx, hy)
        local dir = HEADING(hx, hy, this.x, this.y)
        if not this:isValidHeading() then
            return
        end

        this.dir = dir
    end
    function this:move()
        -- if not this.isValidHeading() then
        --     return
        -- end

        this.x = this.x+(COS(this.dir) * this.speed)
        this.y = this.y+(SIN(this.dir) * this.speed)
    end

    function this:update()
        this.move()
    end

    return this
end

local function makeRects(n)
    for _=0, n do
        table.insert(World, Rect:new())
    end
end

function UPDATE()
    for i, c in pairs(World) do
        c.update()
    end
end

function RENDER()
    -- CLS()
    for _, r in pairs(World) do
        r.render()
    end

    if shouldRenderPallette then
        PALLETTE()
    end
end

function INIT()
    makeRects(nRects)
end
