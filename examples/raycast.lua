local x=0

local Boundaries={}
local Entities={}
local Grid={}
local tileSize=16

local function toTileCoord(x, y)
    return tileSize*x+y
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
function Boundary:new(x0, y0, x1, y1, color)
    local this={x0=x0, y0=y0, x1=x1, y1=y1}
    this.id=UID()
    Boundaries[this.id]=this
    this.color=color

    for i=0, tileSize*x0+x1 do
        Grid[toTileCoord(x0+i, y0)]=this.id
        Grid[toTileCoord(x1+i, y1)]=this.id
    end
    for i=0, tileSize*y0+y1 do
        Grid[toTileCoord(x0, y0+i)]=this.id
        Grid[toTileCoord(x1, y1+i)]=this.id
    end

    function this:render()
        LINE(this.x0, this.y0, this.x1, this.y1, this.color)
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

function is_within_bounds(x, y)
    return x < SCREENW() and x > 0 and y < SCREENH() and y > 0
end



local Ray={}
function Ray:new(x0, y0, dir)
    local this={}
    this.x0=x0
    this.y0=y0
    this.dir=dir
    this.x1=this.x0+COS(dir)*(SCREENW()*2)
    this.y1=this.y0+SIN(dir)*(SCREENH()*2)
    this.color=8
    this.hit_count=0

    function this:is_longer(ox, oy)
        return DISTANCE(this.x0, this.y0, ox, oy) > DISTANCE(this.x0, this.y0, this.x1, this.y1)
    end
    
    function this:reset()
        this.x1=this.x0+COS(this.dir)*SCREENW()
        this.y1=this.y0+SIN(this.dir)*SCREENH()
    end


    -- todo: cast through the grid
    function this:grid_cast()

    end
    -- quick_cast isn't entirely accurate
    --   someof the rays miss because they are reevaluated 
    --   on a different boundary
    function this:quick_cast()
        -- this.color=7
        for _, b in pairs(Boundaries) do
            ox, oy=intersection(
                {x=this.x0, y=this.y0},
                {x=this.x1, y=this.y1},
                {x=b.x0, y=b.y0},
                {x=b.x1, y=b.y1}
            )

            if (ox ~= ox and oy ~= oy) or (ox == nil or oy == nil) then
                return
            end

            local within_line=point_within_line(
                    {x=ox, y=oy}, 
                    {x=b.x0, y=b.y0}, 
                    {x=b.x1, y=b.y1})
            
            local longer=this:is_longer(ox, oy)

            if not is_within_bounds(ox, oy) then
                -- print(ox, oy)
            end

            -- local not_valid= ox ~= ox or oy ~= oy
            -- print(ox, oy)
            if within_line and not longer and is_within_bounds(ox, oy) then
                -- this.color=1
                this.hit_count=this.hit_count+1
                this.x1=ox
                this.y1=oy
            end
        end
        if this.hit_count > 0 then
            this.color=this.hit_count
        end
        
        if this.hit_count == 0 then
            this.color=8
        end
    end
    function this:cast(x0, y0, dir)
        -- reset the ray length
        this:reset()
        this:quick_cast()
    end

    function this:render()
        this:cast(x0, y0)
        LINE(this.x0, this.y0, this.x1, this.y1, this.color)
        this.hit_count=0
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

    -- table.insert(this.rays, Ray:new(this.x, this.y, 16))
    for i=0, 10 do
        table.insert(this.rays, Ray:new(this.x, this.y, i))
    end

    function this:render()
        for _, r in pairs(this.rays) do
            r:render()
        end
        RECT(this.x+(this.r/2), this.y+(this.r/2), this.r*2, this.r*2, 11)
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
        if this.x == SCREENW()/2 or this.y == SCREENH()/2 and NOW() % 3 == 0 then
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
        -- if BTN(0) then -- up
        --     this.y=this.y-1
        -- end
        -- if BTN(1) then -- down
        --     this.y=this.y+1
        -- end
        -- if BTN(2) then -- left
        --     this.x=this.x-1
        -- end
        -- if BTN(3) then -- right
        --     this.x=this.x+1
        -- end
        
        -- if BTN(0) or BTN(1) or BTN(2) or BTN(3) then
        --     return
        -- end

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
    Boundary:new(SCREENW(), 0, 0, 0, 13)
    Boundary:new(SCREENW(), 0, SCREENW(), SCREENH(), 13)
    Boundary:new(0, 0, 0, SCREENH(), 13)
    Boundary:new(0, SCREENH(), SCREENW(), SCREENH(), 13)
    -- Boundary:new(SCREENW(), 2, 0, 2, 13)
    -- Boundary:new(SCREENW()-1, 0, SCREENW()-1, SCREENH(), 13)
    -- Boundary:new(1, 0, 1, SCREENH(), 13)
    -- Boundary:new(0, SCREENH()-2, SCREENW(), SCREENH()-2, 13)

    Boundary:new(60, 60, 80, 80, 13)
    Boundary:new(65, 30, 80, 80, 7)
    Boundary:new(20, 30, 80, 30, 11)
    Boundary:new(80, 30, 100, 50, 7)
    Boundary:new(10, 100, 70, 100, 7)
    table.insert(Entities, Entity:new(100, 75))
    FPS()
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
    RECT(0, SCREENH()-10, 10, 10, 8)
    RECT(10, SCREENH()-10, 10, 10, 1)
    RECT(20, SCREENH()-10, 10, 10, 2)
    RECT(30, SCREENH()-10, 10, 10, 3)
end
