local Boundaries={}
local Entities={}
local Grid={}
local tileSize=16

local function toTileCoord(x, y)
    return tileSize*math.floor(x)+math.floor(y)
end

local CollisionBody={}
function CollisionBody:Create(e)
    local this = {}
    this.x=e.x
    this.y=e.y
    this.w=e.r*2+1
    this.h=e.r*2+1

    function this:isWithinScreen()
        return this.x+this.w < SCREENW() and this.x > 0 and this.y+this.h < SCREENH() and this.y > 0
    end

    function this:isAxisAlignedCollision(cb)
        return this.x < cb.x+cb.w and this.x+ this.w > cb.x and  this.y < cb.y+cb.h and this.h+this.y > cb.y
    end

    return this
end


local Boundary={}
function Boundary:new(x0, y0, x1, y1)
    local this={x0=x0, y0=y0, x1=x1, y1=y1}
    this.id=UID()
    Boundaries[this.id]=this

    for i=0, tileSize*x0+x1 do
        Grid[toTileCoord(x0+i, y0)]=this.id
        Grid[toTileCoord(x1+i, y1)]=this.id
    end
    for i=0, tileSize*y0+y1 do
        Grid[toTileCoord(x0, y0+i)]=this.id
        Grid[toTileCoord(x1, y1+i)]=this.id
    end

    function this:render()
        LINE(this.x0, this.y0, this.x1, this.y1, 7)
    end

    function this:update()
    end

    return this
end

local function intersection (s1, e1, s2, e2)
    local d = (s1.x - e1.x) * (s2.y - e2.y) - (s1.y - e1.y) * (s2.x - e2.x)
    local a = s1.x * e1.y - s1.y * e1.x
    local b = s2.x * e2.y - s2.y * e2.x
    local x = (a * (s2.x - e2.x) - (s1.x - e1.x) * b) / d
    local y = (a * (s2.y - e2.y) - (s1.y - e1.y) * b) / d
    return x, y
end

function point_within_line(p, v0, v1)
    if p.x >= v0.x and p.x <= v1.x and p.y >= v0.y and p.y <= v1.y then
        return true
    end
    return false
end

local Ray={}
function Ray:new(x0, y0, dir)
    local this={}
    this.x0=x0
    this.y0=y0
    this.dir=dir
    this.x1=this.x0+COS(dir)*SCREENW()
    this.y1=this.y0+SIN(dir)*SCREENH()

    function this:is_longer(ox, oy)
        return DISTANCE(this.x0, this.y0, ox, oy) > DISTANCE(this.x0, this.y0, this.x1, this.y1)
    end

    function this:grid_cast()

    end
    -- quick_cast isn't entirely accurate
    --   someof the rays miss
    function this:quick_cast()
        for _, b in pairs(Boundaries) do
            ox, oy=intersection(
                {x=this.x0, y=this.y0},
                {x=this.x1, y=this.y1},
                {x=b.x0, y=b.y0},
                {x=b.x1, y=b.y1}
            )
            if not this:is_longer(ox, oy) and point_within_line({x=ox, y=oy}, {x=b.x0, y=b.y0}, {x=b.x1, y=b.y1}) then
                this.x1=ox
                this.y1=oy
            end
        end
    end
    function this:cast(x0, y0, dir)
        -- reset the ray length
        this.x1=this.x0+COS(this.dir)*SCREENW()
        this.y1=this.y0+SIN(this.dir)*SCREENH()
        this:quick_cast()
    end

    function this:render()
        this:cast(x0, y0)
        LINE(this.x0, this.y0, this.x1, this.y1, 8)
    end
    
    function this:update(x0, y0)
        this.x0=x0
        this.y0=y0
    end

    return this
end

local Entity={}
function Entity:new(x, y)
    local this={x=x, y=y, r=5, speed=1, dir=RND(PI())}
    this.rays={}

    for i=0, 100 do
        table.insert(this.rays, Ray:new(this.x, this.y, i/4))
    end

    function this:render()
        for _, r in pairs(this.rays) do
            r:render()
        end
        CIR(this.x, this.y, this.r, 11)
    end

    function this:isValidHeading()
        if this.dir ~= this.dir then
            return false
        end
        if this.x ~= this.x and this.y ~= this.y and hx ~= hx and hy ~= hy then
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
    function this:random_direction()
        if math.floor(this.x) == SCREENW()/2 or math.floor(this.y) == SCREENH()/2 and NOW() % 3 == 0 then
            this.dir=RND(PI()*2)
        end

        if NOW() % 8 == 0 then
            this.dir=RND(PI()*2)
        end
    end
    function this:move()
        if not CollisionBody:Create(this):isWithinScreen() then
            this.moveToward(64, 64)
        end
        if not this.isValidHeading() then
            return
        end
        this:random_direction()

        this.x = this.x+(COS(this.dir) * this.speed)
        this.y = this.y+(SIN(this.dir) * this.speed)
    end
    function this:update()
        this.move()
        for _, r in pairs(this.rays) do
            r:update(this.x+this.r, this.y+this.r)
        end
    end

    return this
end


function INIT()
    Boundary:new(20, 30, 80, 30)
    Boundary:new(80, 30, 80, 80)
    Boundary:new(60, 60, 80, 80)
    table.insert(Entities, Entity:new(80, 80))
end
function UPDATE()
    for _, b in pairs(Entities) do
        b.update()
    end
end
function RENDER()
    CLS()
    for _, b in pairs(Entities) do
        b.render()
    end
    for _, b in pairs(Boundaries) do
        b.render()
    end
end
