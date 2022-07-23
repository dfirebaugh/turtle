
local World={}
local Grid={}
local size=6
local speed=0.8

local function isAxisAlignedCollision(cb0, cb1)
    return cb0.x < cb1.x+cb1.w and cb0.x+ cb0.w > cb1.x and  cb0.y < cb1.y+cb1.h and cb0.h+cb0.y > cb1.y
end

local CollisionBody={}
function CollisionBody:Create(e)
    this = {}
    this.x=e.x
    this.y=e.y
    this.w=e.size*2+1
    this.h=e.size*2+1

    function this:isWithinScreen()
        return this.x+this.w < SCREENW() and this.x > 0 and this.y+this.h < SCREENH() and this.y > 0
    end

    function this:isAxisAlignedCollision(cb)
        return this.x < cb.x+cb.w and this.x+ this.w > cb.x and  this.y < cb.y+cb.h and this.h+this.y > cb.y
    end

    return this
end

local function toTileCoord(x,y)
    local tileSize=8
    return math.floor(tileSize*math.floor(x)+math.floor(y))
end

local Entity={}
function Entity:new(x, y)
    local this={}
    this.size=RND(size)
    this.x=RND(SCREENW()-this.size)+this.size
    this.y=RND(SCREENH()-this.size)+this.size
    this.speed=speed
    this.color=RND(14)
    this.dir=HEADING(RND(SCREENW()), RND(SCREENH()))
    this.id=UID()
    this.location=toTileCoord(x,y)

    function this:render()
        CIR(this.x, this.y, this.size+1, 1);
        CIR(this.x, this.y, this.size, this.color);
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
    function this:move()
        if not this.isValidHeading() then
            return
        end

        this.x = this.x+(COS(this.dir) * this.speed)
        this.y = this.y+(SIN(this.dir) * this.speed)

        Grid[toTileCoord(this.x,this.y)]=this.id
        Grid[this.location]=0
        this.location=toTileCoord(this.x,this.y)
    end

    function this:checkNeighbors()
        local neighbors = {}

        table.insert(neighbors, Grid[toTileCoord(this.x+1,this.y)])
        table.insert(neighbors, Grid[toTileCoord(this.x-1,this.y)])
        table.insert(neighbors, Grid[toTileCoord(this.x,this.y+1)])
        table.insert(neighbors, Grid[toTileCoord(this.x,this.y-1)])
        table.insert(neighbors, Grid[toTileCoord(this.x+1,this.y+1)])
        table.insert(neighbors, Grid[toTileCoord(this.x-1,this.y-1)])
        table.insert(neighbors, Grid[toTileCoord(this.x+1,this.y-1)])
        table.insert(neighbors, Grid[toTileCoord(this.x-1,this.y+1)])

        for _, n in pairs(neighbors) do
            local e = World[n]
            if n~=0 then
                if isAxisAlignedCollision(CollisionBody:Create(this), CollisionBody:Create(World[n])) then 
                    this.dir = HEADING(this.x, this.y, e.x, e.y)
                end
            end
        --     if n ~= 0 and n ~= nil then
        --     -- if n ~= 0 and n ~= nil and this.x == this.x and this.y == this.y and e.x == e.x and e.y == e.y then
        --         print(n)
        --         if isAxisAlignedCollision(CollisionBody:Create(this), CollisionBody:Create(World[n])) then 
        --             this.dir = HEADING(this.x, this.y, e.x, e.y)
        --         end
        --     end
        end
    end
    function this:update()
        if not CollisionBody:Create(this):isWithinScreen() then
            this.moveToward(64, 64)
        end

        this.checkNeighbors()
        this.move()
    end

    return this
end

local function makeEntities(n)
    for _=0, n do
        -- table.insert(World, Entity:new(RND(SCREENW()), RND(SCREENH())))
        ent = Entity:new(RND(SCREENW()), RND(SCREENH()))
        World[ent.id]=ent
    end
end

function INIT()
    makeEntities(10)
    FPS()
end

count=1
last=0
function UPDATE()
    for _, e in pairs(World) do
        e.update()
    end
    if NOW() ~= last then
        makeEntities(1)
        count=count+1
    end
    last=NOW()
end

function RENDER()
    CLR()
    for _, r in pairs(World) do
        r.render()
    end
    PALLETTE()
    PRINTAT(count, 5, SCREENH()-(SCREENH()*.1), 3)
end
