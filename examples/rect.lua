
local World={}
local size=6
local speed=0.8
local CollisionBody={}

local function isAxisAlignedCollision(cb0, cb1)
    return cb0.x < cb1.x+cb1.w and cb0.x+ cb0.w > cb1.x and  cb0.y < cb1.y+cb1.h and cb0.h+cb0.y > cb1.y
end

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


local Entity={}
function Entity:new()
    local this={}
    this.size=RND(size)
    this.x=RND(SCREENW()-this.size)+this.size
    this.y=RND(SCREENH()-this.size)+this.size
    this.speed=speed
    this.color=RND(14)
    this.dir=HEADING(RND(SCREENW()), RND(SCREENH()))
    this.id=UID()

    function this:render()
        RECT(this.x-1, this.y-1, this.size*2+2, this.size*2+2, 1);
        RECT(this.x, this.y, this.size*2, this.size*2, this.color);
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
    end
    function this:update()
        if not CollisionBody:Create(this):isWithinScreen() then
            this.moveToward(64, 64)
        end
        this.move()
    end

    return this
end

local function makeEntities(n)
    for _=0, n do
        -- table.insert(World, Entity:new(RND(SCREENW()), RND(SCREENH())))
        ent = Entity:new()
        World[ent.id]=ent
    end
end

function INIT()
    makeEntities(400)
end

count=400
last=0
function UPDATE()
    for _, e in pairs(World) do
        e.update()
    end
    if count > 600 then
        return
    end
    if NOW() % 2 == 0 and last ~= NOW() then
        makeEntities(50)
        count=count+10
        -- print(count)
        PRINTAT(count, 0, SCREENH()-10)
    end
    last=NOW()
end

function RENDER()
    CLR()
    for _, r in pairs(World) do
        r.render()
    end
    PALLETTE()
end
